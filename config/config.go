package config

type AppConfig struct {
	ServerPort int
	LogFile    string
}

type DatabaseConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	SslMode            string
	DBName             string
	MaxPoolSize        int
	MaxIdleConnections int
}
