package main

import (
	"fmt"
	"petpujaris/config"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		panic(err)
	}
	config.LoadConfig()
	fmt.Println(config.GetAppConfig())
}
