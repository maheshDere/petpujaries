package models

import (
	"errors"
	"time"
)

type Meals struct {
	ID                  int64     `json:"id" db:"id"`
	Name                string    `json:"name" db:"name"`
	Description         string    `json:"description" db:"description"`
	ImageURL            string    `json:"image_url" db:"image_url"`
	Price               float32   `json:"price" db:"price"`
	Calories            float32   `json:"calories" db:"calories"`
	ISActive            bool      `json:"is_active" db:"is_active"`
	RestaurantCuisineID int64     `json:"restaurant_cuisine_id" db:"restaurant_cuisine_id"`
	MealTypeID          int64     `json:"meal_type_id" db:"meal_type_id"`
	CreatedAt           time.Time `json:"-" db:"created_at"`
	UpdatedAt           time.Time `json:"-" db:"updated_at"`
}

type Ingredients struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type MealsIngredients struct {
	ID           int64     `json:"id" db:"id"`
	MealsID      int64     `json:"meals_id" db:"meals_id"`
	IngredientID int64     `json:"ingredient_id" db:"ingredient_id"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

type Items struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	MealsID   int64     `json:"meals_id" db:"meals_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type MealScheduler struct {
	ID        int64     `db:"id"`
	Date      time.Time `db:"date"`
	MealID    int64     `db:"meal_id"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type MealTypes struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type RestaurantCuisine struct {
	RestaurantCuisineID int64  `db:"id"`
	CuisineName         string `db:"name"`
}

type RestaurantMeal struct {
	MealsID int64  `db:"id"`
	Name    string `db:"name"`
}

func (m Meals) Validation() error {
	if m.Name != "" && m.Description != "" && m.Price != 0 && m.Calories != 0 && m.MealTypeID != 0 && m.RestaurantCuisineID != 0 {
		return nil
	}
	return errors.New("invalid parameter")
}

func (i Items) Validation() error {
	if i.Name != "" && i.MealsID != 0 {
		return nil
	}
	return errors.New("invalid meal item")
}

func (i Ingredients) Validation() error {
	if i.Name != "" {
		return nil
	}
	return errors.New("invalid meal ingredients")
}

func (mi MealsIngredients) Validation() error {
	if mi.MealsID != 0 && mi.IngredientID != 0 {
		return nil
	}
	return errors.New("invalid meal ingredients")
}

func (ms MealScheduler) Validation() error {
	if ms.MealID != 0 && ms.UserID != 0 {
		return nil
	}
	return errors.New("invalid meal scheduler parameter")
}
