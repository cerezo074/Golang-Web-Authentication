package main

import (
	. "example/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/login", loginRoute)
	http.HandleFunc("/logout", logoutRoute)
	http.HandleFunc("/register", registerRoute)
	log.Println("Listening in port 8008...")
	http.ListenAndServe(":8080", nil)
}

func homeRoute(response http.ResponseWriter, request *http.Request) {
	HandleHome(response, request)
}

func loginRoute(response http.ResponseWriter, request *http.Request) {
	HandleLogin(response, request)
}

func logoutRoute(response http.ResponseWriter, request *http.Request) {
	HandleLogout(response, request)
}

func registerRoute(response http.ResponseWriter, request *http.Request) {
	HandleRegister(response, request)
}
