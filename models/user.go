package models

import "time"

type User struct {
	Name             string    `db:"name"`
	Email            string    `db:"email"`
	MobileNumber     string    `db:"mobile_no"`
	IsActive         bool      `db:"is_active"`
	Password         string    `db:"password_digest"`
	RoleID           int       `db:"role_id"`
	ResourceableID   int       `db:"resourceable_id"`
	ResourceableType string    `db:"resourceable_type"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
