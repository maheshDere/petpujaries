package user

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/models"
)

type userService struct {
	registry Registry
}

func (us userService) FindUserByID(ctx context.Context, userID string) (models.User, error) {
	user, err := us.registry.FindUserByID(ctx, userID)
	if err != nil {
		logger.LogError(err, "service.FindUserByID : ", fmt.Sprintf("user not found for id : %v ", userID))
		return models.User{}, err
	}
	return user, nil
}

func NewUserService(registry Registry) Service {
	return userService{registry: registry}
}
