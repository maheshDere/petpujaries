package repository

import "fmt"

var FindUserByID, SaveMealsQuery Command

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
	FindUserByID = Command{
		Table:       "users",
		Description: "fetch user by id",
		Query:       `select id,country_code,mobile_no from %[1]s where id = $1`,
	}

	SaveMealsQuery = Command{
		Table:       "meals",
		Description: "INSERT MEALS DETAILS",
		Query: "insert into %s (name, description, image_url, price, calories, is_active, restaurant_cuisine_id, meal_type_id, created_at,updated_at)" +
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id",
	}
}
