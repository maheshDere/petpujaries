package meals

import (
	"context"
	"petpujaris/models"

	"github.com/stretchr/testify/mock"
)

type MockMealsRegistry struct{ mock.Mock }

func (m *MockMealsRegistry) GetMealType(ctx context.Context) ([]models.MealTypes, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.MealTypes), args.Error(1)
}

func (m *MockMealsRegistry) GetRestaurantCuisine(ctx context.Context, restaurantID int64) ([]models.RestaurantCuisine, error) {
	args := m.Called(ctx, restaurantID)
	return args.Get(0).([]models.RestaurantCuisine), args.Error(1)
}

func (m *MockMealsRegistry) Delete(ctx context.Context, MealID int64) {
	_ = m.Called(ctx, MealID)
	return
}

func (m *MockMealsRegistry) Save(ctx context.Context, mealRecord models.Meals) (int64, error) {
	args := m.Called(ctx, mealRecord)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockMealsRegistry) SaveItem(ctx context.Context, mealItem models.Items) error {
	args := m.Called(ctx, mealItem)
	return args.Error(0)
}
func (m *MockMealsRegistry) SaveIngredients(ctx context.Context, ingredients models.Ingredients) (int64, error) {
	args := m.Called(ctx, ingredients)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMealsRegistry) SaveMealIngredients(ctx context.Context, mealIngredients models.MealsIngredients) error {
	args := m.Called(ctx, mealIngredients)
	return args.Error(0)
}
func (m *MockMealsRegistry) GetRestaurantMeals(ctx context.Context, restaurantID int64) ([]models.RestaurantMeal, error) {
	args := m.Called(ctx, restaurantID)
	return args.Get(0).([]models.RestaurantMeal), args.Error(1)
}
