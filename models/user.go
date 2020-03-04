package models

import "time"

type User struct {
	ID           string    `json:"id" db:"id"`
	CountryCode  string    `json:"country_code" db:"country_code"`
	MobileNumber string    `json:"mobile_number" db:"mobile_no"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}
