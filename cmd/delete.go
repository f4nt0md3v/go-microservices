package cmd

import (
	"github.com/spf13/cobra"
	"go-microservices/cmd/constants"
	nats_api "go-microservices/cmd/nats-api"
	"go-microservices/libs/nats"
)

func InitDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [post_id]",
		Short: "Delete post",
		Long:  `Delete post by post_id`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			go nats.InitSubscribe(nats.CLI_SERVICE, nats_api.Handler, nats_api.ErrorHandler, constants.GetExitChanel())
			constants.NatsRpcCall("DeletePost", map[string]interface{}{
				"id": args[0],
			})

			<-constants.GetExitChanel()
		},
	}

}
