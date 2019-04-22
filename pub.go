package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
)

func main() {
	sc, err := stan.Connect("test-cluster", "dan-pub")
	if err != nil {
		log.Fatalln(err)
	}
	sc.Publish("foo", []byte("Hello"))
	sc.Publish("god", []byte("Hello"))
	sc.Publish("lol", []byte("Hello"))
}
