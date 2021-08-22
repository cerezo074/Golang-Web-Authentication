package controllers

import (
	. "example/persistence"
	"net/http"
)

type JSON map[string]string

func HandleHome(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		sendErrorResponse(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, response)
		return
	}

	user, err := UserForSession(request)

	if err != nil {
		sendErrorResponse("User not authenticated!", http.StatusUnauthorized, response)
		return
	}

	homeResponse := JSON{
		"data": "Welcome " + user.Username + "to your home, we are working hard to give you the best content. Be patient ;)",
	}
	sendSuccessfulResponse(homeResponse, http.StatusOK, response)
}
