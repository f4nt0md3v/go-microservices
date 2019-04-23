# Architecture for HTTP End-Point

- Get `REST HTTP` request to `client-service` for `Post Entity`
- Send message to `Nats Streaming` for any `CRUD`
- `storage-service` get message from `Nats Streaming`
- `storage-service` do any `CRUD` for `Post Entity` 
- `storage-service` send message to `Nats Streaming` 
with any result data (new post, all posts, CRUD status, etc.)
- `client-service` get message from `Nats Streaming`
- `client-service` send response over HTTP with resolve data for client

---
# Architecture for Command Line Interface End-point

- You call command `get`, `new`, `update`, `delete` in terminal
- Send message to `Nats Streaming` for any `CRUD`
- `storage-service` get message from `Nats Streaming`
- `storage-service` do any `CRUD` for `Post Entity` 
- `storage-service` send message to `Nats Streaming` 
with any result data (new post, all posts, CRUD status, etc.)
- Your get message from `Nats Streaming` in terminal

---
# Main Technologies

- [Golang >= 1.12](https://golang.org/)
- [Golang Modules](https://github.com/golang/go/wiki/Modules)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [Nats Streaming](https://nats.io/documentation/streaming/nats-streaming-quickstart/)
- [Protobuf](https://github.com/gogo/protobuf)
- [Cockroach DB](https://www.cockroachlabs.com/)
- [Gorm](http://gorm.io/docs/)
- [Cobra CLI](https://github.com/spf13/cobra)

---
# Build & run in docker-compose

```
$ make clean
$ make
$ cd build
$ docker-compose up -d --force-recreate --build
```

---
# Build Command Line Interface

```
$ C111MODULE=on go build -v cli.go
```

---
# Show

- Use [postman.json](postman.json) for import requests collection 
to [Postman](https://www.getpostman.com/) (`HTTP End-Point`)
- Or use `./cli` for call any commands from your terminal (`CLI End-Point`)