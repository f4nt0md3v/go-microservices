package methods

import (
	"go-microservices/libs/nats"
)

func GetPost(param map[string]interface{}, clientID string) {
	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"title": "!!!",
		},
		Method:    "GetPost",
		ErrorCode: 200,
	})
}

func NewPost(param map[string]interface{}, clientID string) {
	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"success": true,
			"message": "NEW",
		},
		Method: "NewPost",
	})
}

func UpdatePost(param map[string]interface{}, clientID string) {
	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"success": true,
			"message": "UPDATE",
		},
		Method: "UpdatePost",
	})
}

func DeletePost(param map[string]interface{}, clientID string) {
	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"success": true,
			"message": "DELETE",
		},
		Method: "DeletePost",
	})
}
