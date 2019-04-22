package methods

import (
	"go-microservices/services/client-service/helpers"
)

func UniversalNatsRpcReply(param map[string]interface{}, clientID string) {
	helpers.NatsRpcReply(helpers.NatsRpc{
		Param:    param,
		ClientID: clientID,
		IsReply:  true,
	})
}
