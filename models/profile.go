package models

import "time"

type Profile struct {
	ID                   int64     `db:"employee_id"`
	EmployeeID           string    `db:"employee_id"`
	Credits              int64     `db:"credits"`
	UserID               int64     `db:"user_id"`
	NotificationsEnabled bool      `db:"notifications_enabled"`
	MealTypeID           int64     `db:"meal_type_id"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}

func (p *Profile) Validate() []string {
	var errMsg = make([]string, 0)
	if p.EmployeeID == "" {
		errMsg = append(errMsg, "Invalide Employee ID")
	}
	if p.UserID < 0 {
		errMsg = append(errMsg, "Invalide User ID")
	}
	if !(p.MealTypeID == 1 || p.MealTypeID == 2) {
		errMsg = append(errMsg, "Invalide meal type id")
	}
	return errMsg
}
