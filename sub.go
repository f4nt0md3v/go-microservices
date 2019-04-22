package main

import (
	"fmt"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"os"
	"os/signal"
)

func sub(topic string, con stan.Conn) {
	sub, err := con.Subscribe(topic, func(m *stan.Msg) {
		fmt.Printf("Received a message from  %s: %s\n", topic, string(m.Data))
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Unsubscribe()
	select {}
}

func main() {
	sc, err := stan.Connect("test-cluster", "dan-sub")
	if err != nil {
		sc.Close()
		log.Fatalln(err)
	}

	go sub("foo", sc)
	go sub("god", sc)
	go sub("lol", sc)

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Println("Shutdown")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
