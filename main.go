package main

import (
	"net/http"
	"github.homedepot.com/go-api-poc/api/sample"
)

func main() {
	port := "8080"

	router := sample.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	//allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	//allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	http.ListenAndServe(":" + port, router)
	//log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}