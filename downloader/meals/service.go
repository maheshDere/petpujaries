package meals

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/models"
	"petpujaris/repository"
)

type mealsFileService struct {
	mealRegistry repository.MealRegistry
}

type MealsFileService interface {
	GetMealsDetails(ctx context.Context, restaurantID uint64) ([][]string, error)
	GetMealsSchedulerDetails(ctx context.Context, restaurantID uint64) ([][]string, error)
}

func (ms mealsFileService) GetMealsDetails(ctx context.Context, restaurantID uint64) ([][]string, error) {
	mealsDetails := make([][]string, 0)
	mealTypes, err := ms.mealRegistry.GetMealType(ctx)
	if err != nil {
		logger.LogError(err, "downloader.meals", "can not get meal type ")
		return mealsDetails, err
	}
	restaurantCuisines, err := ms.mealRegistry.GetRestaurantCuisine(ctx, int64(restaurantID))
	if err != nil {
		logger.LogError(err, "downloader.meals", "can not get restaurant cuisines")
		return mealsDetails, err
	}

	mealsDetails = createPrimaryMealData(ctx, mealTypes, restaurantCuisines)
	return mealsDetails, nil
}

func createPrimaryMealData(ctx context.Context, mealTypes []models.MealTypes, restaurantCuisines []models.RestaurantCuisine) (mealsDetails [][]string) {

	mealsDetails = make([][]string, 0)

	headers := []string{"MealName", "Description", "ImageUrl", "Price", "Calories", "IsActive", "Item", "Ingredients", "Meal Type Id", "MealType", "Restaurant Cuisine Id", "CuisineName"}
	mealsDetails = append(mealsDetails, headers)

	for i := 0; i < 20; i++ {
		for _, rc := range restaurantCuisines {
			var meals []string
			for _, mt := range mealTypes {
				meals = []string{"", "", "", "", "", "", "", "", fmt.Sprintf("%d", mt.ID), mt.Name, fmt.Sprintf("%d", rc.RestaurantCuisineID), rc.CuisineName}
				mealsDetails = append(mealsDetails, meals)
			}
		}
	}

	return mealsDetails
}

func (ms mealsFileService) GetMealsSchedulerDetails(ctx context.Context, restaurantID uint64) ([][]string, error) {
	mealsSchedulerDetails := make([][]string, 0)

	restaurantMeals, err := ms.mealRegistry.GetRestaurantMeals(ctx, int64(restaurantID))
	if err != nil {
		logger.LogError(err, "downloader.meals", "can not get restaurant meals")
		return mealsSchedulerDetails, err
	}

	mealsSchedulerDetails = createPrimaryMealSchedulerData(ctx, restaurantMeals)
	return mealsSchedulerDetails, nil
}

func createPrimaryMealSchedulerData(ctx context.Context, restaurantMeals []models.RestaurantMeal) (mealsSchedulerDetails [][]string) {

	mealsSchedulerDetails = make([][]string, 0)

	headers := []string{"MealID", "MealName", "Date"}
	mealsSchedulerDetails = append(mealsSchedulerDetails, headers)

	var meals []string
	for _, rm := range restaurantMeals {
		meals = []string{fmt.Sprintf("%d", rm.MealsID), rm.Name, ""}
	}

	mealsSchedulerDetails = append(mealsSchedulerDetails, meals)

	return mealsSchedulerDetails
}
func NewMealsFileService(mr repository.MealRegistry) MealsFileService {
	return mealsFileService{mealRegistry: mr}
}
