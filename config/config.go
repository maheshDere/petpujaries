package config

type AppConfig struct {
	ServerPort int
	LogFile    string `json: required`
}
