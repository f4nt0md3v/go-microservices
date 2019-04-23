.EXPORT_ALL_VARIABLES:

GOOS=linux
GOARCH=amd64
G111MODULE=on
BUILD_DIR=build

.PHONY: \
    all \
    vendor \
    build_client \
    build_storage \
    prepare

all: prepare

vendor:
	go mod vendor

build_client:
	go build -v client-service.go

build_storage:
	go build -v storage-service.go

prepare: build_client build_storage
	mkdir -p $(BUILD_DIR) &&\
	cp nginx.conf $(BUILD_DIR) &&\
	cp Dockerfile.client $(BUILD_DIR) &&\
	cp Dockerfile.storage $(BUILD_DIR) &&\
	cp Dockerfile.nginx $(BUILD_DIR) &&\
	cp docker-compose.yml $(BUILD_DIR) &&\
	cp .env $(BUILD_DIR) &&\
	mv client-service $(BUILD_DIR) &&\
	mv storage-service $(BUILD_DIR)

clean:
	sudo rm -Rf $(BUILD_DIR)