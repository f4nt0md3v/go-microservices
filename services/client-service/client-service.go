// Package implement CLIENT SERVICE.
package client_service

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-microservices/libs/config"
	"go-microservices/libs/errors_handler"
	"go-microservices/libs/logger"
	"go-microservices/libs/nats"
	"go-microservices/services/client-service/constants"
	"go-microservices/services/client-service/helpers"
	nats_api "go-microservices/services/client-service/nats-api"
	"go-microservices/services/client-service/rest"
	"net/http"
	"time"
)

// Create Service.
func CreateService() {
	if !config.GetBool("dev") {
		gin.SetMode(gin.ReleaseMode)
	}

	go nats.InitSubscribe(nats.CLIENT_SREVICE, nats_api.Handler, nats_api.ErrorHandler)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		helpers.SendErrorCall(errors_handler.GetError(100), c)
	})

	rest.InitRest(router)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetInt("client_service_port")),
		Handler:        router,
		ReadTimeout:    constants.TTL * time.Second,
		WriteTimeout:   constants.TTL * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.GetClientService().Info(10)
	err := s.ListenAndServe()
	if err != nil {
		logger.GetClientService().Error(10, map[string]interface{}{
			"error": err,
		})
	}
}
