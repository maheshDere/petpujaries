package repository

import (
	"context"
	"petpujaris/models"
)

type UserRegistry interface {
	Save(ctx context.Context, user models.User) (int64, error)
	GetResourceableID(ctx context.Context, ID uint64) (uint64, error)
	Delete(ctx context.Context, ID int64) (err error)
	SaveProfile(ctx context.Context, profile models.Profile) (err error)
}

type MealRegistry interface {
	Save(ctx context.Context, mealRecord models.Meals) (int64, error)
	SaveItem(ctx context.Context, mealItem models.Items) error
	SaveIngredients(ctx context.Context, ingredients models.Ingredients) (int64, error)
	SaveMealIngredients(ctx context.Context, mealIngredients models.MealsIngredients) error
	Delete(ctx context.Context, MealID int64)
	GetMealType(ctx context.Context) ([]models.MealTypes, error)
	GetRestaurantCuisine(ctx context.Context, restaurantID int64) ([]models.RestaurantCuisine, error)
	GetRestaurantMeals(ctx context.Context, restaurantID int64) ([]models.RestaurantMeal, error)
}

type MealSchedulerRegistry interface {
	Save(ctx context.Context, schedulerRecord models.MealScheduler) error
}
