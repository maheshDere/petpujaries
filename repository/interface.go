package repository

import (
	"context"
	"petpujaris/models"
)

type DatabaseRegistry interface {
	FindUserByID(ctx context.Context, userID string) (models.User, error)
}
