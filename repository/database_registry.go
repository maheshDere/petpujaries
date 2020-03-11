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

func (databaseRegistry databaseRegistry) FindUserByID(ctx context.Context, userID string) (user models.User, err error) {
	err = databaseRegistry.client.QueryRowxContext(ctx, FindUserByID, userID).StructScan(&user)
	if err != nil {
		logger.LogError(err, "FindUserByID", fmt.Sprintf("find user by ID : %v", userID))
		return

	}
	return
}

func (databaseRegistry databaseRegistry) CreateUser(ctx context.Context, user models.User) (err error) {
	_, err = databaseRegistry.client.Exec(ctx, CreateUserQuery, user.Name, user.Email, user.MobileNumber, user.IsActive, user.Password, user.RoleID, user.ResourceableID, user.ResourceableType, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logger.LogError(err, "CreateUser", fmt.Sprintf("create user : %v", user))
		return

	}
	return
}

func NewDBRegistry(pg Client) DatabaseRegistry {
	return databaseRegistry{client: pg}
}
