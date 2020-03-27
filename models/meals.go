package models

import (
	"errors"
	"time"
)

type Meals struct {
	ID                  int64     `db:"id"`
	Name                string    `db:"name"`
	Description         string    `db:"description"`
	ImageURL            string    `db:"image_url"`
	Price               float32   `db:"price"`
	Calories            float32   `db:"calories"`
	ISActive            bool      `db:"is_active"`
	RestaurantCuisineID int64     `db:"restaurant_cuisine_id"`
	MealTypeID          int64     `db:"meal_type_id"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type Ingredients struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type MealsIngredients struct {
	ID           int64     `db:"id"`
	MealsID      int64     `db:"meals_id"`
	IngredientID int64     `db:"ingredient_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Items struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	MealsID   int64     `db:"meals_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
