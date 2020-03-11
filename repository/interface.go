package repository

import (
	"context"
	"petpujaris/models"
)

type UserRegistry interface {
	Save(ctx context.Context, user models.User) (err error)
}

type MealRegistry interface {
	Save(ctx context.Context, mealRecord models.Meals) (int64, error)
	SaveItem(ctx context.Context, mealItem models.Items) error
	SaveIngredients(ctx context.Context, ingredients models.Ingredients) (int64, error)
	SaveMealIngredients(ctx context.Context, mealIngredients models.MealsIngredients) error
}
