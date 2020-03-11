package repository

import "fmt"

var FindUserByID, SaveMealsQuery, SaveMealsItemQuery, SaveIngredientsQuery, SaveMealIngredientsQuery, CreateUserQuery Command

type Command struct {
	Query       string
	Table       string
	Description string
}

func (cmd Command) GetQuery() string {
	return fmt.Sprintf(cmd.Query, cmd.Table)
}

func (cmd Command) GetRawQuery() string {
	return cmd.Query
}

func init() {
	SaveMealsQuery = Command{
		Table:       "meals",
		Description: "INSERT MEALS DETAILS",
		Query: "insert into %s (name, description, image_url, price, calories, is_active, restaurant_cuisine_id, meal_type_id, created_at,updated_at)" +
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id",
	}

	SaveMealsItemQuery = Command{
		Table:       "items",
		Description: "INSERT ITEMS DETAILS",
		Query: "insert into %s (name, meal_id, created_at,updated_at)" +
			" values ($1, $2, $3, $4)",
		//ON CONFLICT (name, meal_id) DO NOTHING  (Need to add indexing on items table create unique INDEX index_name_meal_id on items (name,meal_id);)
	}

	SaveIngredientsQuery = Command{
		Table:       "ingredients",
		Description: "INSERT INGREDIENTS DETAILS",
		Query: "insert into %s (name, created_at,updated_at)" +
			" values ($1, $2, $3) RETURNING id",
		//ON CONFLICT (name) DO NOTHING  (Need to add indexing on ingredients table create unique INDEX index_name on items (name);)
	}

	SaveMealIngredientsQuery = Command{
		Table:       "meal_ingredients",
		Description: "INSERT MEAL INGREDIENTS DETAILS",
		Query: "insert into %s (meal_id,ingredient_id,created_at,updated_at)" +
			" values ($1, $2, $3, $4)",
	}

	CreateUserQuery = Command{
		Table:       "users",
		Description: "INSERT USERS DETAILS",
		Query:       "insert into %s (name, email, mobile_number, is_active, password_digest,  role_id,  resourceable_id,  resourceable_type, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
	}
}
