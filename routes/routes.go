
package routes

import (
"net/http"
"school/controllers"
"school/middleware"
"github.com/gorilla/mux"
)
func SetupRoutes(r *mux.Router) {

	// Test
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // mux 
		w.Write([]byte("API working"))
	}).Methods("GET")

	// Auth controllers 
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Protected Routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Student
	protected.HandleFunc("/students", controllers.AddStudent).Methods("POST")
	protected.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	protected.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	// Teacher
	protected.HandleFunc("/teachers", controllers.AddTeacher).Methods("POST")
	protected.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")

	// Class
	protected.HandleFunc("/class", controllers.AddClass).Methods("POST")
	protected.HandleFunc("/class", controllers.GetClasses).Methods("GET")

	// Fees
	protected.HandleFunc("/fees", controllers.AddFees).Methods("POST")
	protected.HandleFunc("/fees", controllers.GetFees).Methods("GET")

	// Dashboard
	protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")
}