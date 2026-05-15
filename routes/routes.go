
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
	
	// Dashboard admin api full data total students , teacher , 
	protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")

	
	// Attendance
protected.HandleFunc("/attendance", controllers.AddAttendance).Methods("POST")

// Marks
protected.HandleFunc("/marks", controllers.AddMarks).Methods("POST")

// Parent APIs

protected.HandleFunc("/parents", controllers.CreateParent).Methods("POST")

protected.HandleFunc("/parents", controllers.GetParents).Methods("GET")

protected.HandleFunc("/parents/multiple", controllers.AddMultipleParents).Methods("POST")

protected.HandleFunc("/parents/full/{id}", controllers.GetParentFull).Methods("GET")

protected.HandleFunc("/parents/{id}", controllers.UpdateParent).Methods("PUT")

protected.HandleFunc("/parents/{id}", controllers.DeleteParent).Methods("DELETE")

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
    protected.HandleFunc("/teachers", controllers.AddTeacher).Methods("POST")  // add single tecaher 
	protected.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET") // get teacher deatils 
	protected.HandleFunc("/teachers/bulk", controllers.AddMultipleTeachers).Methods("POST")

	protected.HandleFunc("/teachers/{id}", controllers.GetTeacherByID).Methods("GET")
	protected.HandleFunc("/teachers/{id}", controllers.UpdateTeacher).Methods("PUT")
	protected.HandleFunc("/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE")


// SALARY
protected.HandleFunc("/salary", controllers.AddSalary).Methods("POST")
protected.HandleFunc("/salary/{teacherId}", controllers.GetSalaryByTeacher).Methods("GET")

// CLASS
protected.HandleFunc("/classes/bulk", controllers.AddMultipleClasses).Methods("POST")
protected.HandleFunc("/classes", controllers.GetClasses).Methods("GET")



// BULK 
protected.HandleFunc("/subjects/bulk", controllers.AddMultipleSubjects).Methods("POST")


// ASSIGNMENT
protected.HandleFunc("/assignment", controllers.CreateAssignment).Methods("POST")
protected.HandleFunc("/submit", controllers.SubmitAssignment).Methods("POST")

// SUBMISSION
protected.HandleFunc("/submit", controllers.SubmitAssignment).Methods("POST")

// fees
protected.HandleFunc("/create-fee", controllers.CreateFee).Methods("POST")
protected.HandleFunc("/create-bulk-fee", controllers.CreateBulkFee).Methods("POST")
protected.HandleFunc("/pay-fee", controllers.PayFee).Methods("POST")
protected.HandleFunc("/fees", controllers.GetAllFees).Methods("GET")



// transport

r.HandleFunc("/buses", controllers.CreateBus).Methods("POST")
r.HandleFunc("/buses", controllers.GetBuses).Methods("GET")
r.HandleFunc("/buses/{id}", controllers.GetBusByID).Methods("GET")
r.HandleFunc("/buses/bulk", controllers.AddMultipleBuses).Methods("POST")
r.HandleFunc("/transport", controllers.AssignTransport).Methods("POST")
r.HandleFunc("/transport", controllers.GetTransportDetails).Methods("GET")

// STUDENTS
protected.HandleFunc("/students", controllers.AddStudent).Methods("POST")
protected.HandleFunc("/students", controllers.GetStudents).Methods("GET")

protected.HandleFunc("/students/bulk", controllers.AddMultipleStudents).Methods("POST")

protected.HandleFunc("/students/{id}", controllers.GetStudentByID).Methods("GET")

protected.HandleFunc("/students/{id}", controllers.GetStudentByID).Methods("GET")


protected.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")

protected.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")


// LIBRARY
r.HandleFunc("/books", controllers.AddBook).Methods("POST")
r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
r.HandleFunc("/books/bulk", controllers.AddMultipleBooks).Methods("POST")
r.HandleFunc("/library/issue", controllers.IssueBook).Methods("POST")

r.HandleFunc("/library/return", controllers.ReturnBook).Methods("PUT")

r.HandleFunc("/library/details", controllers.GetLibraryDetails).Methods("GET")


// EXAM ROUTES

r.HandleFunc("/exam", controllers.CreateExam).Methods("POST")

r.HandleFunc("/exam/bulk", controllers.AddMultipleExams).Methods("POST")

r.HandleFunc("/exam", controllers.GetExams).Methods("GET")

r.HandleFunc("/exam/{id}", controllers.GetExamByID).Methods("GET")

r.HandleFunc("/exam/{id}", controllers.UpdateExam).Methods("PUT")

r.HandleFunc("/exam/{id}", controllers.DeleteExam).Methods("DELETE")
}