package repository

import (
	"context"
	"petpujaris/config"

	"petpujaris/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgClient struct {
	db *sqlx.DB
}

type Client interface {
	QueryRowxContext(ctx context.Context, cmd Command, args ...interface{}) *sqlx.Row
}

func (pgClient PgClient) QueryRowxContext(ctx context.Context, cmd Command, args ...interface{}) *sqlx.Row {
	return pgClient.db.QueryRowxContext(ctx, cmd.GetQuery(), args...)
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
