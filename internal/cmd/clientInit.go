package cmd

import (
	"github.com/spf13/cobra"
	"webAppCalculator/internal/config"
)

// clientCommand cobra command for making a new client to connect to server
var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "new client",
	Long:  "makes a new client for connecting to server",
	Run: func(cmd *cobra.Command, args []string) {
		config.ServerConfig.RunClient(config.ServerConfig.ConnType,
			config.ServerConfig.ConnHost,
			config.ServerConfig.ConnPort)
	},
}

// init function will initial clientCommand to rootCommand
func init() {
	rootCommand.AddCommand(clientCommand)
}
