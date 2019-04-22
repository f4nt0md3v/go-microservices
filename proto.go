package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"go-microservices/pb"
)

func main() {
	msg := pb.Post{
		OrderId: "!!!",
	}
	var ms pb.Post
	msgM, _ := msg.Marshal()

	proto.Unmarshal(msgM, &ms)
	fmt.Println(msg, msg.Size())
	fmt.Println(ms, ms.Size())
}
