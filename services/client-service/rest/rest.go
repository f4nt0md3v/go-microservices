package rest

import "github.com/gin-gonic/gin"

func InitRest(router *gin.Engine) {
	PostRest(router)
}
