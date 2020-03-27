package models

import (
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

func (m Meals) Validation() []string {
	var errMsg = make([]string, 0)
	if m.Name == "" {
		errMsg = append(errMsg, "Invalid Meal name")
	}
	if m.Description == "" {
		errMsg = append(errMsg, "Invalid Meal Description")
	}
	if m.Price <= 0 {
		errMsg = append(errMsg, "Invalid Meal price")
	}
	if m.Calories <= 0 {
		errMsg = append(errMsg, "Invalid Meal Calories")
	}
	if !(m.MealTypeID == 1 || m.MealTypeID == 2) {
		errMsg = append(errMsg, "Invalid Meal Type ID")
	}
	if m.RestaurantCuisineID == 0 {
		errMsg = append(errMsg, "Invalid Restaurant Cuisine ID")
	}
	return errMsg
}

func (i Items) Validation() []string {
	var errMsg = make([]string, 0)
	if i.Name == "" {
		errMsg = append(errMsg, "Invalid Item")
	}
	if i.MealsID == 0 {
		errMsg = append(errMsg, "Invalid Meal ID")
	}
	return errMsg
}

func (i Ingredients) Validation() []string {
	var errMsg = make([]string, 0)
	if i.Name == "" {
		errMsg = append(errMsg, "Invalid Meal Ingredient")
	}
	return errMsg
}

func (mi MealsIngredients) Validation() []string {
	var errMsg = make([]string, 0)
	if mi.MealsID == 0 {
		errMsg = append(errMsg, "Invalid Meal ID")
	}
	if mi.IngredientID == 0 {
		errMsg = append(errMsg, "Invalid Ingredient ID")

	}
	return errMsg
}

func (ms MealScheduler) Validation() []string {
	var errMsg = make([]string, 0)
	if ms.MealID == 0 {
		errMsg = append(errMsg, "Invalid Meal ID")
	}
	if ms.UserID == 0 {
		errMsg = append(errMsg, "Invalid User ID")
	}
	return errMsg
}
