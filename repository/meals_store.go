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

func NewMealsRegistry(pg Client) MealRegistry {
	return mealsRegistry{client: pg}
}
