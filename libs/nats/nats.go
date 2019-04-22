package nats

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"github.com/satori/go.uuid"
	"go-microservices/libs/config"
	"go-microservices/libs/logger"
)

var conn *stan.Conn

func Connection() *stan.Conn {
	if conn == nil {
		uid := uuid.NewV4()
		sc, err := stan.Connect(config.GetString("nats_cluster"), uid.String())
		if err != nil {
			logger.GetNats().Error(10)
			sc.Close()
		}
		conn = &sc
	}
	return conn
}

func Publish(topic string, msg proto.Message) error {
	s, err := proto.Marshal(msg)
	if err != nil {
		logger.GetNats().Error(30, map[string]interface{}{
			"topic": topic,
		})
		return err
	}

	err = (*Connection()).Publish(topic, s)
	if err != nil {
		logger.GetNats().Error(40, map[string]interface{}{
			"error": err,
			"topic": topic,
		})
		return err
	}

	if config.GetBool("dev") {
		logger.GetNats().Info(20, map[string]interface{}{
			"topic":   topic,
			"message": msg,
		})
	}
	return nil
}

func InitSubscribe(topic string, handler func(msg proto.Message)) {
	con := Connection()
	if conn == nil {
		return
	}

	sub, err := (*con).Subscribe(topic, func(m *stan.Msg) {

		var message proto.Message
		err := proto.Unmarshal(m.Data, message)
		if err != nil {
			logger.GetNats().Error(30, map[string]interface{}{
				"topic": topic,
			})
			return
		}

		if config.GetBool("dev") {
			logger.GetNats().Info(30, map[string]interface{}{
				"message": message,
			})
		}

		if handler != nil {
			handler(*message)
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		logger.GetNats().Error(20)
	} else {
		logger.GetNats().Info(10, map[string]interface{}{
			"topic": topic,
		})
		defer sub.Unsubscribe()
		select {}
	}

}
