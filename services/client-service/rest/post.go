package rest

import (
	"github.com/gin-gonic/gin"
	"go-microservices/libs/errors_handler"
	"go-microservices/libs/nats"
	"go-microservices/services/client-service/helpers"
)

func PostRest(router *gin.RouterGroup) {
	router.GET("/posts", getPosts)
	router.GET("/post/:id", getPost)
	router.POST("/post", newPost)
	router.PUT("/post/:id", updatePost)
	router.DELETE("/post/:id", deletePost)
}

type Post struct {
	Title string
}

func getPosts(c *gin.Context) {
	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "GetPosts",
		IsReply: true,
		Context: c,
	})
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(19, map[string]interface{}{
			"field":   "id",
			"message": "Field must be int",
		}), c)
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
	post := Post{}
	err := c.BindJSON(&post)
	if err != nil {
		helpers.SendErrorCall(errors_handler.GetError(11), c)
		return
	}
	if post.Title == "" {
		helpers.SendErrorCall(errors_handler.GetError(17, map[string]interface{}{
			"field": "title",
		}), c)
		return
	}

	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "NewPost",
		Param: map[string]interface{}{
			"title": post.Title,
		},
		IsReply: true,
		Context: c,
	})
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(19, map[string]interface{}{
			"field":   "id",
			"message": "Field must be int",
		}), c)
		return
	}

	post := Post{}
	err := c.BindJSON(&post)
	if err != nil {
		helpers.SendErrorCall(errors_handler.GetError(11), c)
		return
	}
	if post.Title == "" {
		helpers.SendErrorCall(errors_handler.GetError(18), c)
		return
	}

	helpers.NatsRpcCall(helpers.NatsRpc{
		Service: nats.STORAGE_SERVICE,
		Method:  "UpdatePost",
		Param: map[string]interface{}{
			"id":    id,
			"title": post.Title,
		},
		IsReply: true,
		Context: c,
	})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helpers.SendErrorCall(errors_handler.GetError(19, map[string]interface{}{
			"field":   "id",
			"message": "Field must be int",
		}), c)
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
