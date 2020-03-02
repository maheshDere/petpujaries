package config

import "fmt"

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

func (dbcfg DatabaseConfig) DataSourceURL() string {
	/* db://user:secret@localhost:6379/0?foo=bar&qux=baz */
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbcfg.User, dbcfg.Password, dbcfg.Host, dbcfg.Port, dbcfg.DBName, dbcfg.SslMode)
}
