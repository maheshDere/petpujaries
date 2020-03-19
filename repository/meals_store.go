package repository

import (
	"context"
	"petpujaris/models"
)

type mealsRegistry struct {
	client Client
}

func (mr mealsRegistry) Save(ctx context.Context, mealRecord models.Meals) (int64, error) {
	var id int64
	rows, err := mr.client.Query(ctx, SaveMealsQuery, mealRecord.Name, mealRecord.Description,
		mealRecord.ImageURL, mealRecord.Price, mealRecord.Calories, mealRecord.ISActive,
		mealRecord.RestaurantCuisineID, mealRecord.MealTypeID, mealRecord.CreatedAt, mealRecord.UpdatedAt)
	if err != nil {
		return id, err
	}

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return id, err
		}
	}

	return id, err
}

func (mr mealsRegistry) SaveItem(ctx context.Context, mealItem models.Items) error {
	_, err := mr.client.Exec(ctx, SaveMealsItemQuery, mealItem.Name, mealItem.MealsID, mealItem.CreatedAt, mealItem.UpdatedAt)
	return err

}

func (mr mealsRegistry) SaveIngredients(ctx context.Context, ingredients models.Ingredients) (int64, error) {
	var id int64
	rows, err := mr.client.Query(ctx, SaveIngredientsQuery, ingredients.Name, ingredients.CreatedAt, ingredients.UpdatedAt)
	if err != nil {
		return id, err
	}

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return id, err
		}
	}

	return id, err
}

func (mr mealsRegistry) SaveMealIngredients(ctx context.Context, mealIngredients models.MealsIngredients) error {
	_, err := mr.client.Exec(ctx, SaveMealIngredientsQuery, mealIngredients.MealsID, mealIngredients.IngredientID, mealIngredients.CreatedAt, mealIngredients.UpdatedAt)
	return err
}

func (mr mealsRegistry) Delete(ctx context.Context, MealID int64) {
	mr.client.Exec(ctx, DeleteMealIngredientsQuery, MealID)
	mr.client.Exec(ctx, DeleteMealsItemQuery, MealID)
	mr.client.Exec(ctx, DeleteMealsQuery, MealID)

}

func (mr mealsRegistry) GetMealType(ctx context.Context) ([]models.MealTypes, error) {
	var mealTypes []models.MealTypes
	rows, err := mr.client.Query(ctx, GetMealTypeQuery)
	if err != nil {
		return mealTypes, err
	}

	defer rows.Close()
	for rows.Next() {
		var mealType models.MealTypes
		if err := rows.Scan(&mealType.ID, &mealType.Name, &mealType.CreatedAt, &mealType.UpdatedAt); err != nil {
			return mealTypes, err
		}
		mealTypes = append(mealTypes, mealType)
	}

	return mealTypes, nil
}

func (mr mealsRegistry) GetRestaurantCuisine(ctx context.Context, restaurantID int64) ([]models.RestaurantCuisine, error) {
	var restaurantCuisines []models.RestaurantCuisine
	rows, err := mr.client.Query(ctx, GetRestaurantCuisineQuery, restaurantID)
	if err != nil {
		return restaurantCuisines, err
	}

	defer rows.Close()
	for rows.Next() {
		var restaurantCuisine models.RestaurantCuisine
		if err := rows.Scan(&restaurantCuisine.RestaurantCuisineID, &restaurantCuisine.CuisineName); err != nil {
			return restaurantCuisines, err
		}
		restaurantCuisines = append(restaurantCuisines, restaurantCuisine)
	}

	return restaurantCuisines, nil
}

func (mr mealsRegistry) GetRestaurantMeals(ctx context.Context, restaurantID int64) ([]models.RestaurantMeal, error) {
	var restaurantMeals []models.RestaurantMeal
	rows, err := mr.client.Query(ctx, GetRestaurantMealQuery, restaurantID)
	if err != nil {
		return restaurantMeals, err
	}

	defer rows.Close()
	for rows.Next() {
		var restaurantMeal models.RestaurantMeal
		if err := rows.Scan(&restaurantMeal.MealsID, &restaurantMeal.Name); err != nil {
			return restaurantMeals, err
		}
		restaurantMeals = append(restaurantMeals, restaurantMeal)
	}

	return restaurantMeals, nil
}
func NewMealsRegistry(pg Client) MealRegistry {
	return mealsRegistry{client: pg}
}
