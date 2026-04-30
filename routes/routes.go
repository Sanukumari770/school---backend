
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
	
	protected.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	protected.HandleFunc("/student", controllers.CreateStudent).Methods("POST")
	protected.HandleFunc("/student/{id}", controllers.GetStudentFull).Methods("GET")
	protected.HandleFunc("/student/{id}", controllers.UpdateStudent).Methods("PUT")
	protected.HandleFunc("/student/{id}", controllers.DeleteStudent).Methods("DELETE")
	

	// Teacher
protected.HandleFunc("/teacher", controllers.AddTeacher).Methods("POST")
protected.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")
protected.HandleFunc("/teacher/{id}", controllers.GetTeacherFull).Methods("GET")
protected.HandleFunc("/teacher/{id}", controllers.UpdateTeacher).Methods("PUT")
protected.HandleFunc("/teacher/{id}", controllers.DeleteTeacher).Methods("DELETE")
	// Class
protected.HandleFunc("/class", controllers.AddClass).Methods("POST")
protected.HandleFunc("/class", controllers.GetClasses).Methods("GET")

	// Fees
protected.HandleFunc("/fee/create", controllers.CreateFee).Methods("POST")
protected.HandleFunc("/fee/pay", controllers.PayFee).Methods("POST")
protected.HandleFunc("/fee/{id}", controllers.GetFeeDetails).Methods("GET")

	// Dashboard
	protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")

	// Admission
	protected.HandleFunc("/admission/apply", controllers.ApplyAdmission).Methods("POST")
	protected.HandleFunc("/admission/{id}/entrance", controllers.UpdateEntrance).Methods("PUT")
	protected.HandleFunc("/admission/{id}/approve", controllers.ApproveAdmission).Methods("POST")
	protected.HandleFunc("/admission/{id}", controllers.GetAdmissionFull).Methods("GET")

	// Attendance
protected.HandleFunc("/attendance", controllers.AddAttendance).Methods("POST")


// Exam
protected.HandleFunc("/exam", controllers.CreateExam).Methods("POST")

// Marks
protected.HandleFunc("/marks", controllers.AddMarks).Methods("POST")
}