package repository

import (
	"petpujaris/config"

	"petpujaris/logger"

	"github.com/jmoiron/sqlx"
)

type PgClient struct {
	db *sqlx.DB
}

func NewPgClient(dbcfg config.DatabaseConfig) (PgClient, error) {
	db, err := GetDBConnection(dbcfg)
	if err != nil {
		logger.LogError(err, "postgres new client", "create new connection failure")
		return PgClient{}, err
	}

	return PgClient{db}, nil
}

func GetDBConnection(dbcfg config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbcfg.DataSourceURL())
	if err != nil {
		logger.LogError(err, "postgres client", "Opening connection failure")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.LogError(err, "postgres client", "DB Ping failure")
		return nil, err
	}

	db.SetMaxIdleConns(dbcfg.MaxIdleConnections)
	return db, nil
}
