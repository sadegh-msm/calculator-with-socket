package main

import webAppCalculator "webAppCalculator/internal"

func main() {
	server := webAppCalculator.Server{
		ConnHost: "localhost",
		ConnPort: "8080",
		ConnType: "tcp",
	}
	server.RunClient(server.ConnType, server.ConnHost, server.ConnType)
}
