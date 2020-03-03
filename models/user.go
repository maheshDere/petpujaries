package models

import "time"

type User struct {
	ID           string    `db:"id"`
	CountryCode  string    `db:"country_code"`
	MobileNumber string    `db:"mobile_no"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
