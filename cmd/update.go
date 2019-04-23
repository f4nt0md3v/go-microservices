package cmd

import (
	"github.com/spf13/cobra"
	"go-microservices/cmd/constants"
	nats_api "go-microservices/cmd/nats-api"
	"go-microservices/libs/nats"
)

func InitUpdate() *cobra.Command {
	return &cobra.Command{
		Use:   "update [post_id] [title]",
		Short: "Update post",
		Long:  `Update post title by post_id`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			go nats.InitSubscribe(nats.CLI_SERVICE, nats_api.Handler, nats_api.ErrorHandler, constants.GetExitChanel())
			constants.NatsRpcCall("UpdatePost", map[string]interface{}{
				"id":    args[0],
				"title": args[1],
			})

			<-constants.GetExitChanel()
		},
	}

}
