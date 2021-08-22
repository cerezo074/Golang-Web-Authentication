package models

import (
	"fmt"
)

type User struct {
	Username string
	Email    string
	Password string
}

func (element User) IsValidID() bool {
	return element.Email != "" &&
		element.Email != " " &&
		len(element.Email) > 0
}

func (element User) Validate() error {
	if !isValid(element.Email) {
		return fmt.Errorf("invalid email field, value: %v", element.Email)
	}

	if !isValid(element.Password) {
		return fmt.Errorf("invalid password field, value: %v", element.Password)
	}

	if !isValid(element.Username) {
		return fmt.Errorf("invalid username field, value: %v", element.Username)
	}

	return nil
}

func isValid(value string) bool {
	return value != "" &&
		value != " " &&
		len(value) > 0
}
