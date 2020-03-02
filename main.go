package main

import (
	"petpujaris/config"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		panic(err)
	}
	config.LoadConfig()

}
