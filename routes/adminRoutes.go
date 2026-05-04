package routes

import (
	"school/controllers"
"github.com/gorilla/mux"
)
func AdminRoutes(router *mux.Router) {

	router.HandleFunc("/admin/student", controllers.AddMultipleStudents).Methods("POST")
	router.HandleFunc("/admin/class", controllers.CreateClass).Methods("POST")
	router.HandleFunc("/admin/subject", controllers.CreateSubject).Methods("POST")
}