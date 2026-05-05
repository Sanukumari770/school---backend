
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
	protected.HandleFunc("/students/bulk", controllers.AddMultipleStudents).Methods("POST")
	protected.HandleFunc("/student/{id}", controllers.GetStudentFull).Methods("GET")
	protected.HandleFunc("/student/{id}", controllers.UpdateStudent).Methods("PUT")
	protected.HandleFunc("/student/{id}", controllers.DeleteStudent).Methods("DELETE")
	
	// Fees
protected.HandleFunc("/fee/create", controllers.CreateFee).Methods("POST")
protected.HandleFunc("/fee/pay", controllers.PayFee).Methods("POST")
protected.HandleFunc("/fee/{id}", controllers.GetFeeDetails).Methods("GET")

	// Dashboard admin api full data total students , teacher , 
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

// Parent APIs
protected.HandleFunc("/parent", controllers.CreateParent).Methods("POST")
protected.HandleFunc("/parents", controllers.GetParents).Methods("GET")
protected.Handle(
	"/parent/{id}",
	middleware.Authorize("parent")(http.HandlerFunc(controllers.GetParentFull)),
).Methods("GET")

protected.HandleFunc("/parent/{id}", controllers.UpdateParent).Methods("PUT")
protected.HandleFunc("/parent/{id}", controllers.DeleteParent).Methods("DELETE")



// Parent protected
parentRouter := r.PathPrefix("/parent").Subrouter()
parentRouter.Use(middleware.AuthMiddleware)
parentRouter.Use(middleware.Authorize("parent"))
parentRouter.HandleFunc("/dashboard/{id}", controllers.GetParentDashboard).Methods("GET")

//------------ Role based acces control  middleware ------------------
// Admin only
protected.Handle("/teacher", middleware.Authorize("admin")(http.HandlerFunc(controllers.AddTeacher))).Methods("POST")

// Teacher only
protected.Handle("/attendance", middleware.Authorize("teacher")(http.HandlerFunc(controllers.AddAttendance))).Methods("POST")

// Parent only
protected.Handle("/parent/{id}", middleware.Authorize("parent")(http.HandlerFunc(controllers.GetParentFull))).Methods("GET")

// dashboard parents 
protected.Handle(
	"/parent/dashboard",
	middleware.Authorize("parent")(http.HandlerFunc(controllers.GetParentDashboard)),
).Methods("GET")
// TEACHER MODULE
protected.Handle("/teacher",
		middleware.Authorize("admin")(http.HandlerFunc(controllers.AddTeacher)),
	).Methods("POST")

	protected.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")
	protected.HandleFunc("/teachers/bulk", controllers.AddMultipleTeachers).Methods("POST")

	protected.HandleFunc("/teacher/{id}", controllers.GetTeacherFull).Methods("GET")
	protected.HandleFunc("/teacher/{id}", controllers.UpdateTeacher).Methods("PUT")
	protected.HandleFunc("/teacher/{id}", controllers.DeleteTeacher).Methods("DELETE")


// SALARY
protected.HandleFunc("/salary", controllers.AddSalary).Methods("POST")
protected.HandleFunc("/salary/{teacherId}", controllers.GetSalaryByTeacher).Methods("GET")

// CLASS
protected.HandleFunc("/classes/bulk", controllers.AddMultipleClasses).Methods("POST")
protected.HandleFunc("/classes", controllers.GetClasses).Methods("GET")



// ASSIGNMENT
protected.HandleFunc("/assignment", controllers.CreateAssignment).Methods("POST")
protected.HandleFunc("/submit", controllers.SubmitAssignment).Methods("POST")

// SUBMISSION
protected.HandleFunc("/submit", controllers.SubmitAssignment).Methods("POST")
}