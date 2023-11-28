package main

import (
	"cryptom.app/config"
	"cryptom.app/server"
)

func main() {
	config.Load()

	serverStarted := make(chan bool)

	server := &server.Server{
		Port:            config.ServerPort(),
		MongoDbUri:      config.MongoDbConnectionUrl(),
		RpcUrl:          config.RpcUrl(),
		AllowedOrigins:  config.AllowedOrigins(),
		ServerReady:     serverStarted,
		StorageHostname: config.StorageHostname(),
	}

	server.Start()
}
