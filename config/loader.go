package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var appConfig AppConfig
var dbConfig DatabaseConfig

func SetupConfig() error {
	configFile := "application"

	environment := os.Getenv("ENVIRONMENT")
	if environment == "test" {
		configFile = "application-test"
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./../")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("SetupConfig: Config file not found")
		}
		return fmt.Errorf("SetupConfig: Config file can not load errors %v", err)
	}
	return nil
}

func LoadConfig() {
	var errs []error
	if err := LoadAppConfig(); err != nil {
		errs = append(errs, err)
	}

	if err := LoadDBConfig(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		panic(errs)
	}

}

func LoadAppConfig() error {
	if err := viper.Unmarshal(&appConfig); err != nil {
		return fmt.Errorf("LoadAppConfig : fail to load app config, %v", err)
	}
	return nil
}

func LoadDBConfig() error {
	if err := viper.Unmarshal(&dbConfig); err != nil {
		return fmt.Errorf("LoadDBConfig : fail to load db config, %v", err)
	}
	return nil
}

func GetAppConfig() AppConfig {
	return appConfig
}

func GetDBConfig() DatabaseConfig {
	return dbConfig
}
