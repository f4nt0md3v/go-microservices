package cmd

import (
	"github.com/spf13/cobra"
	"go-microservices/cmd/constants"
	nats_api "go-microservices/cmd/nats-api"
	"go-microservices/libs/nats"
)

func InitNew() *cobra.Command {
	return &cobra.Command{
		Use:   "new [title]",
		Short: "New post",
		Long:  `Create new post with title`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			go nats.InitSubscribe(nats.CLI_SERVICE, nats_api.Handler, nats_api.ErrorHandler, constants.GetExitChanel())
			constants.NatsRpcCall("NewPost", map[string]interface{}{
				"title": args[0],
			})

			<-constants.GetExitChanel()
		},
	}

}
