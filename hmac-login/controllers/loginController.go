package controllers

import (
	. "example/persistence"
	"net/http"
)

func HandleLogin(response http.ResponseWriter, request *http.Request) {
	user, _ := UserForSession(request)

	if user != nil {
		sendErrorResponse("User authenticated! previously", http.StatusForbidden, response)
		return
	}

	email := request.Form.Get("email")
	user, err := GetUser(email)
	if err != nil {
		sendErrorResponse("User not found", http.StatusNotFound, response)
		return
	}

	password := request.Form.Get("password")
	if !IsPasswordValid(password, user.Email) {
		sendErrorResponse("Invalid password", http.StatusNotFound, response)
		return
	}

	newToken, err := RegisterToken(*user)
	if err != nil {
		sendErrorResponse("Invalid session", http.StatusInternalServerError, response)
		return
	}

	successfulJSON := JSON{"token": "Bearer " + newToken}
	sendSuccessfulResponse(successfulJSON, http.StatusOK, response)
}

func HandleLogout(response http.ResponseWriter, request *http.Request) {
	user, err := UserForSession(request)

	if err != nil {
		sendErrorResponse("User not authenticated!", http.StatusUnauthorized, response)
		return
	}

	err = DeleteToken(*user)
	if err != nil {
		sendErrorResponse("Invalid session", http.StatusUnauthorized, response)
		return
	}

	successfulJSON := JSON{"data": "BYE"}
	sendSuccessfulResponse(successfulJSON, http.StatusOK, response)
}
