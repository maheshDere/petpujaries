package repository

import (
	"context"
	"petpujaris/models"
)

type DatabaseRegistry interface {
	FindUserByID(ctx context.Context, userID string) (models.User, error)
}

type MealRegistry interface {
	Save(ctx context.Context, mealRecord models.Meals) (int64, error)
	SaveItem(ctx context.Context, mealItem models.Items) error
	SaveIngredients(ctx context.Context, ingredients models.Ingredients) (int64, error)
	SaveMealIngredients(ctx context.Context, mealIngredients models.MealsIngredients) error
}
