package models

import (
	"fmt"
	"petpujaris/logger"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var MOBILE_NUMBER_REGEX = regexp.MustCompile("^(?:(?:\\+|0{0,2})91(\\s*[\\-]\\s*)?|[0]?)\\d{10}$")
var USER_NAME_REGEX = regexp.MustCompile("^[a-zA-Z]+(?:[\\s.]+[a-zA-Z]+)*$")

type User struct {
	Name             string    `db:"name"`
	Email            string    `db:"email"`
	MobileNumber     string    `db:"mobile_no"`
	IsActive         bool      `db:"is_active"`
	Password         string    `db:"password_digest"`
	RoleID           int       `db:"role_id"`
	ResourceableID   uint64    `db:"resourceable_id"`
	ResourceableType string    `db:"resourceable_type"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	Profile          Profile
}

func (u *User) GenerateHashedPassword(key string) (hashedPassword string, err error) {
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(key), 10)
	if err != nil {
		logger.LogError(err, "model.User.GenerateHashedPassword", fmt.Sprintf("error to generate hashPassword for user %v with password  %v ", u.Name, u.Password))
		return
	}

	return string(byteHashPassword), nil
}

func (u *User) Validate() []string {
	var errMsg = make([]string, 0)
	if !isValidUserName(u.Name) {
		errMsg = append(errMsg, "Invalid user name")
	}
	if u.Email == "" {
		errMsg = append(errMsg, "Invalid email")
	}
	if !isValidMobileNumber(u.MobileNumber) {
		errMsg = append(errMsg, "Invalid mobile number")
	}
	if u.RoleID < 0 {
		errMsg = append(errMsg, "Invalid role ID")
	}
	if u.ResourceableID < 0 {
		errMsg = append(errMsg, "Invalid Resourceable ID")
	}
	if u.ResourceableType == "" {
		errMsg = append(errMsg, "Invalid Resourceable type")

	}
	return errMsg
}

func isValidMobileNumber(mobileNumber string) bool {
	if MOBILE_NUMBER_REGEX.MatchString(mobileNumber) {
		return true
	}
	return false
}

func isValidUserName(name string) bool {
	if USER_NAME_REGEX.MatchString(name) {
		return true
	}
	return false
}
