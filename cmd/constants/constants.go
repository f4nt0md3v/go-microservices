package constants

import (
	uuid "github.com/satori/go.uuid"
	"go-microservices/libs/nats"
	"sync"
)

var exist_chan chan interface{} = nil

func GetExitChanel() chan interface{} {
	if exist_chan == nil {
		exist_chan = make(chan interface{})
	}
	return exist_chan
}

var clientID string = ""
var clientMutexR = sync.RWMutex{}
var clientMutexW = sync.RWMutex{}

func NatsRpcCall(method string, details map[string]interface{}) {
	cl := uuid.NewV4().String()
	SetClientID(cl)
	nats.Publish(nats.NatsMessage{
		Services:  []string{nats.STORAGE_SERVICE},
		Method:    method,
		IsSuccess: true,
		Details:   details,
		ClientID:  cl,
	})
}

func GetClientID() string {
	clientMutexR.Lock()
	defer clientMutexR.Unlock()
	return clientID
}

func SetClientID(s string) {
	clientMutexW.Lock()
	defer clientMutexW.Unlock()
	clientID = s
}
