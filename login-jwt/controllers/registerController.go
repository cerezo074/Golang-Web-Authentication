package controllers

import (
	. "example/models"
	. "example/persistence"
	"fmt"
	"net/http"
)

func HandleRegister(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		sendErrorResponse(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, response)
		return
	}

	email := request.FormValue("email")
	username := request.FormValue("username")
	password := request.FormValue("password")

	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := user.Validate()
	if err != nil {
		sendErrorResponse(fmt.Sprintf("Error, %v", err), http.StatusBadRequest, response)
		return
	}

	err = Register(username, email, password)
	if err != nil {
		message := fmt.Sprintf("Cant create user, %v", err)
		sendErrorResponse(message, http.StatusBadRequest, response)
		return
	}

	successfulJSON := JSON{"data": "user created!"}
	sendSuccessfulResponse(successfulJSON, http.StatusOK, response)
}
