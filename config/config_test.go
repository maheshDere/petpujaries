package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseSourceURL(t *testing.T) {
	databaseConfig := getDatabaseConfig()

	expectedSourceURL := "postgres://SampleUser:SamplePassword@SampleHost:5432/SampleDbName?sslmode=disable"
	actualSourceURL := databaseConfig.DataSourceURL()

	assert.Equal(t, expectedSourceURL, actualSourceURL)
}

func TestGetDBConfig(t *testing.T) {
	err := SetupConfig()
	LoadConfig()
	assert.NoError(t, err)
	dbconfig := GetDBConfig()
	assert.NotEmpty(t, dbconfig)
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:               "SampleHost",
		Port:               5432,
		User:               "SampleUser",
		Password:           "SamplePassword",
		SslMode:            "disable",
		DBName:             "SampleDbName",
		MaxPoolSize:        10,
		MaxIdleConnections: 5,
	}
}
