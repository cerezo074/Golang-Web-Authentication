package controllers

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(info string, status int, response http.ResponseWriter) {
	message := JSON{"error": info}
	jsonData, err := json.Marshal(message)
	dispathJSON(jsonData, status, err, response)
}

func sendSuccessfulResponse(info JSON, status int, response http.ResponseWriter) {
	jsonData, err := json.Marshal(info)
	dispathJSON(jsonData, status, err, response)
}

func dispathJSON(data []byte, status int, err error, response http.ResponseWriter) {
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	response.Write(data)
}
