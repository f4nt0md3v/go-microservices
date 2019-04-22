// Package implement any common HTTP helpers.
package helpers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-microservices/libs/errors_handler"
	"go-microservices/libs/nats"
	"go-microservices/services/client-service/constants"
	"go-microservices/services/client-service/storage"
	"net/http"
)

// Send encapsulation object for SUCCESS JSON.
func SendSuccessCall(result interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, result)
}

// Send resolve ERROR format object for ERROR JSON.
func SendErrorCall(err errors_handler.Error, c *gin.Context) {
	if err.Details() != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    err.Code(),
			"message": err.Error(),
			"details": err.Details(),
		})
	} else {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    err.Code(),
			"message": err.Error(),
		})
	}
}

type NatsRpc struct {
	Service  string
	Method   string
	Param    map[string]interface{}
	ClientID string
	IsReply  bool
	Context  *gin.Context
}

func NatsRpcCall(n NatsRpc) (bool, map[string]interface{}, int) {
	ch := make(chan constants.InternalMessage)

	uid := uuid.NewV4().String()
	storage.GetStorage().AddToStorage(uid, ch)
	nats.Publish(nats.NatsMessage{
		Service:   n.Service,
		IsSuccess: true,
		Method:    n.Method,
		Details:   n.Param,
		ClientID:  uid,
	})

	msg := <-ch
	close(ch)
	if n.IsReply {
		if msg.IsSuccess {
			SendSuccessCall(msg.Details, n.Context)
		} else {
			SendErrorCall(errors_handler.GetError(int(msg.ErrorCode), msg.Details), n.Context)
		}
	}

	return msg.IsSuccess, msg.Details, int(msg.ErrorCode)
}

func NatsRpcReply(p NatsRpc) map[string]interface{} {
	ch := storage.GetStorage().GetChan(p.ClientID)
	if ch == nil {
		return nil
	}

	ch <- constants.InternalMessage{
		IsSuccess: true,
		Details:   p.Param,
	}

	storage.GetStorage().RemoveFromStorage(p.ClientID)
	return p.Param
}
