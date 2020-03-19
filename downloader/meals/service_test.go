package meals

import (
	"context"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	logger.Setup()
}

func TestGetMealsDetails(t *testing.T) {
	t.Run("when pass valid restaurant ID", func(t *testing.T) {
		restaurantID := 2
		ctx := context.Background()
		mr, mfs := setup()

		mr.On("GetMealType", ctx).Return([]models.MealTypes{models.MealTypes{int64(1), "Veg", time.Now(), time.Now()}, models.MealTypes{int64(1), "Non Veg", time.Now(), time.Now()}}, nil).Once()
		mr.On("GetRestaurantCuisine", ctx, int64(restaurantID)).Return([]models.RestaurantCuisine{models.RestaurantCuisine{int64(2), "CHINESE"}}, nil).Once()
		t.Run("It should return primary meals details", func(t *testing.T) {
			mealsPrimaryDetials, err := mfs.GetMealsDetails(ctx, uint64(restaurantID))
			assert.NoError(t, err)
			assert.NotEmpty(t, mealsPrimaryDetials)
			mr.AssertExpectations(t)
		})
	})

}

func TestGetMealsSchedulerDetails(t *testing.T) {
	t.Run("when pass valid restaurant ID", func(t *testing.T) {
		restaurantID := 2
		ctx := context.Background()
		mr, mfs := setup()

		mr.On("GetRestaurantMeals", ctx, int64(restaurantID)).Return([]models.RestaurantMeal{models.RestaurantMeal{int64(2), "PANEER MASALA"}}, nil).Once()
		t.Run("It should return primary meals details", func(t *testing.T) {
			mealsPrimaryDetials, err := mfs.GetMealsSchedulerDetails(ctx, uint64(restaurantID))
			assert.NoError(t, err)
			assert.NotEmpty(t, mealsPrimaryDetials)
			mr.AssertExpectations(t)
		})
	})

}

func setup() (*MockMealsRegistry, MealsFileService) {
	mr := new(MockMealsRegistry)
	return mr, NewMealsFileService(mr)
}
