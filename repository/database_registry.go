package repository

import (
	"context"
	"fmt"
	"petpujaris/logger"

	"petpujaris/models"
)

type databaseRegistry struct {
	client Client
}

type DatabaseRegistry interface {
	FindUserByID(ctx context.Context, userID string) (models.User, error)
}

func (databaseRegistry databaseRegistry) FindUserByID(ctx context.Context, userID string) (user models.User, err error) {
	err = databaseRegistry.client.QueryRowxContext(ctx, FindUserByID, userID).StructScan(&user)
	if err != nil {
		logger.LogError(err, "FindUserByID", fmt.Sprintf("find user by ID : %v", userID))
		return

	}
	return
}

func NewDBRegistry(pg Client) DatabaseRegistry {
	return databaseRegistry{client: pg}
}
