package repository

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/models"
)

type userRegistry struct {
	client Client
}

const RollID = 3

func (ur userRegistry) Save(ctx context.Context, user models.User) (err error) {
	_, err = ur.client.Exec(ctx, CreateUserQuery, user.Name, user.Email, user.MobileNumber, user.IsActive, user.Password, user.RoleID, user.ResourceableID, user.ResourceableType, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logger.LogError(err, "CreateUser", fmt.Sprintf("create user : %v", user))
		return
	}

	return
}

func (ur userRegistry) GetResourceableID(ctx context.Context, ID uint64) (uint64, error) {
	var resourceableID uint64
	err := ur.client.QueryRow(ctx, GetResourceableIDQuery, ID, RollID).Scan(&resourceableID)
	if err != nil {
		return resourceableID, err
	}
	return resourceableID, nil
}

func NewUserRegistry(pg Client) userRegistry {
	return userRegistry{client: pg}
}
