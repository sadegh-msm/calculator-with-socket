package config

import webAppCalculator "webAppCalculator/internal"

// ServerConfig is a configuration for sever stats
var ServerConfig = webAppCalculator.Server{
	ConnType: "tcp",
	ConnHost: "localhost",
	ConnPort: "8080",
}
