package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var appConfig AppConfig

func SetupConfig() error {
	configFile := "application"

	environment := os.Getenv("ENVIRONMENT")
	if environment == "test" {
		configFile = "application-test"
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("SetupConfig: Config file not found")
		}
		return fmt.Errorf("SetupConfig: Config file can not load errors", err)
	}
	return nil
}

func LoadConfig() {
	var errs []error
	if err := LoadAppConfig(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		panic(errs)
	}

}

func LoadAppConfig() error {
	if err := viper.Unmarshal(&appConfig); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	return nil
}

func GetAppConfig() AppConfig {
	return appConfig
}
