package cmd

import (
	"github.com/spf13/cobra"
	webAppCalculator "webAppCalculator/internal"
)

var (
	server = webAppCalculator.Server{
		ConnHost: "localhost",
		ConnPort: "8080",
		ConnType: "tcp",
	}

	StartServer = &cobra.Command{
		Use:   "start",
		Short: "starts the server for clients to connect",
		Run: func(cmd *cobra.Command, args []string) {
			server.RunServer(server.ConnType, server.ConnHost, server.ConnPort)
		},
	}

	NewClient = &cobra.Command{
		Use:   "new",
		Short: "builds a new client to connect to server",
		Run: func(cmd *cobra.Command, args []string) {
			server.RunClient(server.ConnType, server.ConnHost, server.ConnPort)
		},
	}
)

func Init() {
	StartServer.AddCommand()
	NewClient.AddCommand()
}
