package cmd

import (
	"github.com/spf13/cobra"
	"webAppCalculator/internal/config"
)

// serverCommand cobra command for starting a server with cmd
var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "starts server",
	Long:  "this command starts the server for clients to connect it",
	Run: func(cmd *cobra.Command, args []string) {
		config.ServerConfig.RunClient(config.ServerConfig.ConnType,
			config.ServerConfig.ConnHost,
			config.ServerConfig.ConnPort)
	},
}

// init function will initial serverCommand to rootCommand
func init() {
	rootCommand.AddCommand(serverCommand)
}
