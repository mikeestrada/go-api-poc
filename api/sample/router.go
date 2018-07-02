package sample

import (
	"github.com/gorilla/mux"
)

var sampleApi = &SampleApi{}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/sample", sampleApi.SampleApiHandler).Methods("GET")

	sampleApi.DBConnection = "sample connection"
	router.HandleFunc("/sample/configured", sampleApi.SampleHandler()).Methods("GET")
	return router
}