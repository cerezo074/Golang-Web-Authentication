package models

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

func (element User) IsValid() bool {
	return isValid(element.Email) && isValid(element.Password) && isValid(element.Username)
}

func isValid(value string) bool {
	return value != "" &&
		value != " " &&
		len(value) > 0
}
