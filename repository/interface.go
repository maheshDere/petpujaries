package repository

import (
	"context"
	"petpujaris/models"
)

type DatabaseRegistry interface {
	FindUserByID(ctx context.Context, userID string) (models.User, error)
}

type MealRegistry interface {
	Save(ctx context.Context, mealRecord models.Meals) (int8, error)
}
