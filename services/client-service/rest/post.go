package rest

import (
	"github.com/gin-gonic/gin"
	"go-microservices/libs/errors_handler"
	"go-microservices/libs/nats"
	"go-microservices/services/client-service/helpers"
)

func PostRest(router *gin.Engine) {
	router.GET("/post/:id", getPost)
	router.POST("/post", newPost)
	router.PUT("/post/:id", updatePost)
	router.DELETE("/post/:id", deletePost)
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(200), c)
		return
	}

	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "GetPost",
		Param: map[string]interface{}{
			"id": id,
		},
		IsReply: true,
		Context: c,
	})
}

func newPost(c *gin.Context) {
	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "NewPost",
		Param: map[string]interface{}{
			"title": "twesdzsd d",
		},
		IsReply: true,
		Context: c,
	})
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(200), c)
		return
	}

	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "UpdatePost",
		Param: map[string]interface{}{
			"title": "twesdzsd d",
		},
		IsReply: true,
		Context: c,
	})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(200), c)
		return
	}

	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "DeletePost",
		Param: map[string]interface{}{
			"id": id,
		},
		IsReply: true,
		Context: c,
	})
}
