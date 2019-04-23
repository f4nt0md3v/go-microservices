package nats

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	nats_controller "github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/satori/go.uuid"
	"go-microservices/libs/config"
	"go-microservices/libs/json_codec"
	"go-microservices/libs/logger"
	"go-microservices/pb"
	"os"
	"sync"
)

var conn *stan.Conn
var connMutexR = sync.RWMutex{}

func Connection() *stan.Conn {
	connMutexR.Lock()
	defer connMutexR.Unlock()
	if conn == nil {
		uid := uuid.NewV4()

		nc, err := nats_controller.Connect(fmt.Sprintf("nats://%s:%d", config.GetString("nats_host"), config.GetInt("nats_port")),
			nats_controller.UserInfo(config.GetString("nats_user"), config.GetString("nats_pass")))
		if err != nil {
			logger.GetNats().Error(10)
		}

		sc, err := stan.Connect(config.GetString("nats_cluster"), uid.String(), stan.NatsConn(nc))
		if err != nil {
			logger.GetNats().Error(10)
			sc.Close()
		}
		conn = &sc
	}
	return conn
}

type NatsMessage struct {
	Services  []string
	ErrorCode int32
	IsSuccess bool
	Method    string
	Details   map[string]interface{}
	ClientID  string
}

func Publish(m NatsMessage) {
	details, encodeError := json_codec.JsonEncode(m.Details)
	if encodeError != nil {
		return
	}

	s, err := proto.Marshal(&pb.Nats{
		IsSuccess:  m.IsSuccess,
		ErrorCode:  m.ErrorCode,
		Method:     m.Method,
		ClientUuid: m.ClientID,
		Details:    details,
	})
	if err != nil {
		logger.GetNats().Error(30, map[string]interface{}{
			"services": m.Services,
		})
		return
	}
	for _, service := range m.Services {
		go func(serviceName string) {
			_, err = (*Connection()).PublishAsync(serviceName, s, nil)
			if err != nil {
				logger.GetNats().Error(40, map[string]interface{}{
					"error":   err,
					"service": serviceName,
					"message": m,
				})
				return
			}

			if config.GetBool("dev") {
				logger.GetNats().Info(20, map[string]interface{}{
					"service": serviceName,
					"message": m,
				})
			}
		}(service)
	}
}

func InitSubscribe(topic string, handler map[string]func(map[string]interface{}, string), errorHandle func(int32, map[string]interface{}, string), close chan interface{}) {
	con := Connection()
	if conn == nil {
		os.Exit(1)
	}

	sub, err := (*con).Subscribe(topic, func(m *stan.Msg) {
		var message pb.Nats
		err := proto.Unmarshal(m.Data, &message)
		if err != nil {
			logger.GetNats().Error(30)
			return
		}

		if config.GetBool("dev") {
			logger.GetNats().Info(30, map[string]interface{}{
				"message": message,
			})
		}

		if handler != nil {
			if message.GetIsSuccess() {
				method, okMethod := handler[message.GetMethod()]
				details, err := json_codec.JsonParse(message.GetDetails())
				if okMethod {
					if err == nil {
						method(details, message.GetClientUuid())
					} else {
						method(nil, message.GetClientUuid())
					}
				}
			} else {
				details, err := json_codec.JsonParse(message.GetDetails())
				if err == nil {
					errorHandle(message.GetErrorCode(), details, message.GetClientUuid())
				}
			}
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		logger.GetNats().Error(20)
	} else {
		logger.GetNats().Info(10, map[string]interface{}{
			"services": topic,
		})
		defer sub.Unsubscribe()
		if close != nil {
			<-close
		} else {
			select {}
		}
	}
}
