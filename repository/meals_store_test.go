package repository

import (
	"context"
	"fmt"
	"petpujaris/config"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mealRegistry MealRegistry

func init() {
	config.SetupConfig()
	logger.Setup()
	config.LoadConfig()
	dbConfig := config.GetDBConfig()
	fmt.Println(dbConfig)
	pgClient, err := NewPgClient(dbConfig)
	if err != nil {
		panic(err)
	}

	mealRegistry = NewMealsRegistry(pgClient)
}
func TestMealsRegistry_Save(t *testing.T) {
	t.Run("when insert valid meals data", func(t *testing.T) {
		mealRecord := models.Meals{
			Name:                "Paneer Masala",
			Description:         "Mix Multiple Content",
			ImageURL:            "NA",
			MealTypeID:          10,
			Calories:            100,
			Price:               10,
			ISActive:            true,
			RestaurantCuisineID: 2,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}

		id, err := mealRegistry.Save(context.Background(), mealRecord)
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEqual(t, int64(0), id)
		})
	})
}

func TestMealsRegistry_SaveItem(t *testing.T) {
	t.Run("when insert valid meals item", func(t *testing.T) {
		mealItemRecord := models.Items{
			Name:      "Paneer",
			MealsID:   10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := mealRegistry.SaveItem(context.Background(), mealItemRecord)
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
		})
	})
}

func TestMealsRegistry_SaveIngredients(t *testing.T) {
	t.Run("when insert valid meals ingredients", func(t *testing.T) {
		IngredientRecord := models.Ingredients{
			Name:      "Paneer",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		id, err := mealRegistry.SaveIngredients(context.Background(), IngredientRecord)
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEqual(t, int64(0), id)
		})
	})
}

func TestMealsRegistry_SaveMealsIngredients(t *testing.T) {
	t.Run("when insert valid meals ingredients relation", func(t *testing.T) {
		mealIngredientsRecord := models.MealsIngredients{
			MealsID:      10,
			IngredientID: 12,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		err := mealRegistry.SaveMealIngredients(context.Background(), mealIngredientsRecord)
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
		})
	})
}

func TestMealsRegistry_GetMealType(t *testing.T) {
	t.Run("Get MealsType records", func(t *testing.T) {
		mealsTypes, err := mealRegistry.GetMealType(context.Background())
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEmpty(t, mealsTypes)
		})
	})
}

func TestMealsRegistry_GetRestaurantCuisine(t *testing.T) {
	t.Run("Get MealsType records", func(t *testing.T) {
		restaurantID := 2
		restaurantCuisine, err := mealRegistry.GetRestaurantCuisine(context.Background(), int64(restaurantID))
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEmpty(t, restaurantCuisine)
		})
	})
}
