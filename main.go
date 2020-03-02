package main

import (
	"errors"
	"fmt"
	"petpujaris/config"
	"petpujaris/logger"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		panic(err)
	}

	config.LoadConfig()
	logger.Setup()
	logger.LogError(errors.New("err"), "main", "can not run script")
	fmt.Println(config.GetAppConfig())

}
