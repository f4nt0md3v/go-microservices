// Package implement STORAGE SERVICE.
package storage_service

import (
	"go-microservices/libs/cockroach"
	"go-microservices/libs/nats"
	nats_api "go-microservices/services/storage-service/nats-api"
)

// Create Service.
func CreateService() {
	cockroach.Connection()
	//defer cockroach.Connection().clo
	go nats.InitSubscribe(nats.STORAGE_SERVICE, nats_api.Handler, nats_api.ErrorHandler)
	select {}
}
