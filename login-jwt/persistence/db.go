package persistence

import (
	. "example/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type store map[string]User

func (element store) contains(email string) bool {
	return element[email].IsValidID()
}

var (
	users *store = &store{}
)

func Register(username string, email string, password string) error {
	if users.contains(email) {
		return fmt.Errorf("user with email %s exists", email)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	(*users)[email] = User{Username: username, Email: email, Password: string(bytes)}
	return nil
}

func GetUser(email string) (*User, error) {
	if !users.contains(email) {
		return nil, fmt.Errorf("user with the following %s email doesn't exist", email)
	}

	selectedUser := (*users)[email]

	return &selectedUser, nil
}

func IsPasswordValid(password string, email string) bool {
	if !users.contains(email) {
		return false
	}

	selectedUser := (*users)[email]

	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(password))
	return err == nil
}
