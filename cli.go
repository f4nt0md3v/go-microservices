package main

import (
	"github.com/spf13/cobra"
	"go-microservices/cmd"
	"go-microservices/libs/logger"
)

func main() {
	logger.SetIsLogOut(false)
	var rootCmd = &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cmd.InitGet())
	rootCmd.AddCommand(cmd.InitNew())
	rootCmd.AddCommand(cmd.InitUpdate())
	rootCmd.AddCommand(cmd.InitDelete())
	rootCmd.Execute()
}
