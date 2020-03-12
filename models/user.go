package models

import (
	"fmt"
	"petpujaris/logger"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

func (u *User) GenerateHashedPassword() (hashedPassword string, err error) {
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		logger.LogError(err, "model.User.GenerateHashedPassword", fmt.Sprintf("error to generate hashPassword for user %v with password", u.Name, u.Password))
		return
	}

	return string(byteHashPassword), nil
}
