package user

import (
	"context"
	"petpujaris/models"
)

type Service interface {
	FindUserByID(context.Context, string) (models.User, error)
}

type Registry interface {
	FindUserByID(ctx context.Context, userID string) (models.User, error)
}
