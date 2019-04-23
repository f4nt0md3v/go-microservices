package nats_api

import (
	"go-microservices/services/storage-service/nats-api/methods"
)

var Handler map[string]func(map[string]interface{}, string) = map[string]func(map[string]interface{}, string){
	"GetPosts":   methods.GetPosts,
	"GetPost":    methods.GetPost,
	"NewPost":    methods.NewPost,
	"UpdatePost": methods.UpdatePost,
	"DeletePost": methods.DeletePost,
}

func ErrorHandler(code int32, details map[string]interface{}, clientID string) {

}
