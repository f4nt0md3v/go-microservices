package methods

import (
	"fmt"
	"go-microservices/cmd/constants"
	"gopkg.in/yaml.v2"
	"os"
)

func UniversalNatsRpcReply(param map[string]interface{}, clientID string) {
	if clientID == constants.GetClientID() {
		r, err := yaml.Marshal(param)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Print(string(r))
		close(constants.GetExitChanel())
	}
}
