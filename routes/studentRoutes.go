package routes

import (
	"school/controllers"

	"github.com/gorilla/mux"
)

func StudentRoutes(router *mux.Router) {

	router.HandleFunc("/student/{id}", controllers.GetStudentFull).Methods("GET")
}