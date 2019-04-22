package nats_api

import (
	"go-microservices/services/client-service/constants"
	"go-microservices/services/client-service/nats-api/methods"
	"go-microservices/services/client-service/storage"
)

var Handler map[string]func(map[string]interface{}, string) = map[string]func(map[string]interface{}, string){
	"GetPost":    methods.UniversalNatsRpcReply,
	"NewPost":    methods.UniversalNatsRpcReply,
	"UpdatePost": methods.UniversalNatsRpcReply,
	"DeletePost": methods.UniversalNatsRpcReply,
}

func ErrorHandler(code int32, details map[string]interface{}, clientID string) {
	if clientID != "" {
		ch := storage.GetStorage().GetChan(clientID)
		if ch == nil {
			return
		}

		ch <- constants.InternalMessage{
			IsSuccess: false,
			ErrorCode: code,
			Details:   details,
		}
		storage.GetStorage().RemoveFromStorage(clientID)
	}
}
