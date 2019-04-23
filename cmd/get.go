package cmd

import (
	"github.com/spf13/cobra"
	"go-microservices/cmd/constants"
	nats_api "go-microservices/cmd/nats-api"
	"go-microservices/libs/nats"
)

func InitGet() *cobra.Command {
	return &cobra.Command{
		Use:   "get [post_id|all]",
		Short: "Get post or posts",
		Long:  `Get information about single post or all posts`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			go nats.InitSubscribe(nats.CLI_SERVICE, nats_api.Handler, nats_api.ErrorHandler, constants.GetExitChanel())

			if args[0] == "all" {
				constants.NatsRpcCall("GetPosts", nil)
			} else {
				constants.NatsRpcCall("GetPost", map[string]interface{}{
					"id": args[0],
				})
			}

			<-constants.GetExitChanel()
		},
	}

}
