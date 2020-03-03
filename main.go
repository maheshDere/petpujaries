package main

import (
	"petpujaris/config"
	"petpujaris/logger"
	"petpujaris/server"
)

func main() {
	err := startServer()
	panic(err)
}

func startServer() error {
	if err := config.SetupConfig(); err != nil {
		panic(err)
	}

	config.LoadConfig()
	logger.Setup()
	defer logger.Close()

	server := server.HTTP{Port: config.ServerPort()}
	err := server.Start()
	if err != nil {
		logger.LogError(err, "main.server.Start", "error in start server")
		return err
	}
	return nil
}
