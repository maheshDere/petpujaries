package models

import "time"

type Meals struct {
	ID                  int8      `json:"id" db:"id"`
	Description         string    `json:"description" db:"description"`
	ImageURL            string    `json:"image_url" db:"image_url"`
	Price               float32   `json:"price" db:"price"`
	Calories            float32   `json:"calories" db:"calories"`
	ISActive            bool      `json:"is_active" db:"is_active"`
	RestaurantCuisineID int8      `json:"restaurant_cuisine_id" db:"restaurant_cuisine_id"`
	MealTypeID          int8      `json:"meal_type_id" db:"meal_type_id"`
	CreatedAt           time.Time `json:"-" db:"created_at"`
	UpdatedAt           time.Time `json:"-" db:"updated_at"`
}

type Ingredients struct {
	ID        int8      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type MealsIngredients struct {
	ID           int8      `json:"id" db:"id"`
	MealsID      int8      `json:"meals_id" db:"meals_id"`
	IngredientID int8      `json:"ingredient_id" db:"ingredient_id"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

type Items struct {
	ID        int8      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	MealsID   int8      `json:"meals_id" db:"meals_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
