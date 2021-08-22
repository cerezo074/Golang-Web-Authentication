package controllers

import (
	. "example/models"
	. "example/persistence"
	"fmt"
	"net/http"
)

func HandleRegister(response http.ResponseWriter, request *http.Request) {
	email := request.Form.Get("email")
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if !user.IsValid() {
		sendErrorResponse("Invalid fields, please fill all!", http.StatusBadRequest, response)
		return
	}

	err := Register(username, email, password)

	if err != nil {
		message := fmt.Sprintf("Cant create user, %v", err)
		sendErrorResponse(message, http.StatusBadRequest, response)
		return
	}

	successfulJSON := JSON{"data": "user created!"}
	sendSuccessfulResponse(successfulJSON, http.StatusOK, response)
}
