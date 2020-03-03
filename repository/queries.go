package repository

import "fmt"

var FindUserByID Command

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
}
