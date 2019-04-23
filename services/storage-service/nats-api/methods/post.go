package methods

import (
	"github.com/thoas/go-funk"
	"go-microservices/libs/cockroach"
	"go-microservices/libs/json_codec"
	"go-microservices/libs/nats"
	"strconv"
)

func GetPosts(param map[string]interface{}, clientID string) {
	posts, err := cockroach.GetAllPosts()
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "GetPosts",
			ErrorCode: 10,
		})
		return
	}

	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"posts": funk.Map(posts, func(post cockroach.Post) map[string]interface{} {
				return map[string]interface{}{
					"id":    strconv.FormatInt(post.ID, 10),
					"title": post.Title,
					"date":  post.CreatedAt.String(),
				}
			}),
		},
		Method: "GetPosts",
	})
}

func GetPost(param map[string]interface{}, clientID string) {
	id, err := getID("DeletePost", param, clientID)
	if err != nil {
		return
	}

	post, err := cockroach.GetPost(id)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "GetPost",
			ErrorCode: 10,
		})
		return
	}

	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"id":    strconv.FormatInt(post.ID, 10),
			"title": post.Title,
			"date":  post.CreatedAt.String(),
		},
		Method: "GetPost",
	})
}

func NewPost(param map[string]interface{}, clientID string) {
	title, err := json_codec.GetString("title", param)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "NewPost",
			ErrorCode: 10,
		})
		return
	}

	post, err := cockroach.CreatePost(cockroach.Post{
		Title: title,
	})
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "NewPost",
			ErrorCode: 10,
		})
		return
	}

	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"id":    strconv.FormatInt(post.ID, 10),
			"title": post.Title,
			"date":  post.CreatedAt.String(),
		},
		Method: "NewPost",
	})
}

func UpdatePost(param map[string]interface{}, clientID string) {
	title, err := json_codec.GetString("title", param)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "UpdatePost",
			ErrorCode: 10,
		})
		return
	}
	id, err := getID("DeletePost", param, clientID)
	if err != nil {
		return
	}

	err = cockroach.UpdatePost(cockroach.Post{ID: id}, cockroach.Post{Title: title})
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "UpdatePost",
			ErrorCode: 10,
		})
		return
	}

	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"success": true,
		},
		Method: "UpdatePost",
	})
}

func DeletePost(param map[string]interface{}, clientID string) {
	id, err := getID("DeletePost", param, clientID)
	if err != nil {
		return
	}

	err = cockroach.DeletePost(id)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    "DeletePost",
			ErrorCode: 10,
		})
		return
	}

	nats.Publish(nats.NatsMessage{
		Service:   nats.CLIENT_SREVICE,
		ClientID:  clientID,
		IsSuccess: true,
		Details: map[string]interface{}{
			"success": true,
		},
		Method: "DeletePost",
	})
}

func getID(method string, param map[string]interface{}, clientID string) (int64, error) {
	i, err := json_codec.GetString("id", param)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    method,
			ErrorCode: 18,
			Details: map[string]interface{}{
				"field": "id",
			},
		})
		return 0, err
	}
	id, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		nats.Publish(nats.NatsMessage{
			Service:   nats.CLIENT_SREVICE,
			ClientID:  clientID,
			IsSuccess: false,
			Method:    method,
			ErrorCode: 20,
			Details: map[string]interface{}{
				"field": "id",
				"type":  "int",
			},
		})
		return 0, err
	}
	return id, nil
}
