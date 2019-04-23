package nats_api

import (
	"fmt"
	"go-microservices/cmd/constants"
	"go-microservices/cmd/nats-api/methods"
	"go-microservices/libs/errors_handler"
	"gopkg.in/yaml.v2"
	"os"
)

var Handler map[string]func(map[string]interface{}, string) = map[string]func(map[string]interface{}, string){
	"GetPosts":   methods.UniversalNatsRpcReply,
	"GetPost":    methods.UniversalNatsRpcReply,
	"NewPost":    methods.UniversalNatsRpcReply,
	"UpdatePost": methods.UniversalNatsRpcReply,
	"DeletePost": methods.UniversalNatsRpcReply,
}

func ErrorHandler(code int32, details map[string]interface{}, clientID string) {
	if clientID == constants.GetClientID() {
		r, err := yaml.Marshal(errors_handler.GetError(int(code), details))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Error!\n\n")
		fmt.Print(string(r))
		close(constants.GetExitChanel())
	}
}
