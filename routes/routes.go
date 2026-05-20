package routes

import (
	"net/http"

	auth "school/controllers/auth"
	"school/controllers"
	"school/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {

	// TEST ROUTE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API working"))
	}).Methods("GET")

	// ================= AUTH ROUTES =================

	r.HandleFunc("/register", controllers.Register).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/student/login", auth.StudentLoginController).Methods("POST")

	r.HandleFunc("/teacher/login", auth.TeacherLoginController).Methods("POST")

	r.HandleFunc("/parent/login", auth.ParentLoginController).Methods("POST")

	// ================= PUBLIC NOTICE ROUTES =================

	r.HandleFunc("/notices", controllers.CreateNotice).Methods("POST")

	r.HandleFunc("/notices/bulk", controllers.AddMultipleNotices).Methods("POST")

	r.HandleFunc("/notices", controllers.GetNotices).Methods("GET")

	r.HandleFunc("/notices/{id}", controllers.GetNotice).Methods("GET")

	r.HandleFunc("/notices/{id}", controllers.DeleteNotice).Methods("DELETE")

	// ================= PROTECTED ROUTES =================

	protected := r.PathPrefix("/api").Subrouter()

	protected.Use(middleware.AuthMiddleware)

	// ================= ATTENDANCE =================

	protected.HandleFunc("/attendance", controllers.AddAttendance).Methods("POST")

	// ================= MARKS =================

	protected.HandleFunc("/marks", controllers.AddMarks).Methods("POST")

	// ================= PARENTS =================

	protected.HandleFunc("/parents", controllers.CreateParent).Methods("POST")

	protected.HandleFunc("/parents", controllers.GetParents).Methods("GET")

	protected.HandleFunc("/parents/multiple", controllers.AddMultipleParents).Methods("POST")

	protected.HandleFunc("/parents/full/{id}", controllers.GetParentFull).Methods("GET")

	protected.HandleFunc("/parents/{id}", controllers.UpdateParent).Methods("PUT")

	protected.HandleFunc("/parents/{id}", controllers.DeleteParent).Methods("DELETE")

	// ================= ROLE BASED ACCESS =================

	// Admin only
	protected.Handle(
		"/teacher",
		middleware.Authorize("admin")(http.HandlerFunc(controllers.AddTeacher)),
	).Methods("POST")

	// Teacher only
	protected.Handle(
		"/attendance",
		middleware.Authorize("teacher")(http.HandlerFunc(controllers.AddAttendance)),
	).Methods("POST")

	// Parent only
	protected.Handle(
		"/parent/{id}",
		middleware.Authorize("parent")(http.HandlerFunc(controllers.GetParentFull)),
	).Methods("GET")

	// ================= PARENT DASHBOARD =================

	parentRouter := protected.PathPrefix("/parent").Subrouter()

	parentRouter.Use(middleware.Authorize("parent"))

	parentRouter.HandleFunc(
		"/dashboard/{id}",
		controllers.GetParentDashboard,
	).Methods("GET")

	// ================= TEACHERS =================

	r.HandleFunc("/teachers", controllers.AddTeacher).Methods("POST")

	r.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")

	r.HandleFunc("/teachers/bulk", controllers.AddMultipleTeachers).Methods("POST")

	r.HandleFunc("/teachers/{id}", controllers.GetTeacherByID).Methods("GET")

	r.HandleFunc("/teachers/{id}", controllers.UpdateTeacher).Methods("PUT")

	r.HandleFunc("/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE")

	// ================= SALARY =================

	protected.HandleFunc("/salary", controllers.AddSalary).Methods("POST")

	protected.HandleFunc("/salary/{teacherId}", controllers.GetSalaryByTeacher).Methods("GET")

	// ================= CLASSES =================

	protected.HandleFunc("/classes/bulk", controllers.AddMultipleClasses).Methods("POST")

	protected.HandleFunc("/classes", controllers.GetClasses).Methods("GET")

	// ================= SUBJECTS =================

	protected.HandleFunc("/subjects/bulk", controllers.AddMultipleSubjects).Methods("POST")

	// ================= FEES =================

	protected.HandleFunc("/create-fee", controllers.CreateFee).Methods("POST")

	protected.HandleFunc("/create-bulk-fee", controllers.CreateBulkFee).Methods("POST")

	protected.HandleFunc("/pay-fee", controllers.PayFee).Methods("POST")

	protected.HandleFunc("/fees", controllers.GetAllFees).Methods("GET")

	// ================= STUDENTS =================

	r.HandleFunc("/students", controllers.AddStudent).Methods("POST")

	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")

	r.HandleFunc("/students/bulk", controllers.AddMultipleStudents).Methods("POST")

	r.HandleFunc("/students/{id}", controllers.GetStudentByID).Methods("GET")

	r.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")

	r.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	// ================= STUDENT DASHBOARD =================

	protected.HandleFunc(
		"/student/dashboard",
		controllers.GetStudentDashboard,
	).Methods("GET")

	// ================= TRANSPORT =================

	r.HandleFunc("/buses", controllers.CreateBus).Methods("POST")

	r.HandleFunc("/buses", controllers.GetBuses).Methods("GET")

	r.HandleFunc("/buses/{id}", controllers.GetBusByID).Methods("GET")

	r.HandleFunc("/buses/bulk", controllers.AddMultipleBuses).Methods("POST")

	r.HandleFunc("/transport", controllers.AssignTransport).Methods("POST")

	r.HandleFunc("/transport", controllers.GetTransportDetails).Methods("GET")

	// ================= LIBRARY =================

	r.HandleFunc("/books", controllers.AddBook).Methods("POST")

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")

	r.HandleFunc("/books/bulk", controllers.AddMultipleBooks).Methods("POST")

	r.HandleFunc("/library/issue", controllers.IssueBook).Methods("POST")

	r.HandleFunc("/library/return", controllers.ReturnBook).Methods("PUT")

	r.HandleFunc("/library/details", controllers.GetLibraryDetails).Methods("GET")

	// ================= EXAMS =================

	r.HandleFunc("/exam", controllers.CreateExam).Methods("POST")

	r.HandleFunc("/exam/bulk", controllers.AddMultipleExams).Methods("POST")

	r.HandleFunc("/exam", controllers.GetExams).Methods("GET")

	r.HandleFunc("/exam/{id}", controllers.GetExamByID).Methods("GET")

	r.HandleFunc("/exam/{id}", controllers.UpdateExam).Methods("PUT")

	r.HandleFunc("/exam/{id}", controllers.DeleteExam).Methods("DELETE")

	// ================= ASSIGNMENTS =================

	r.HandleFunc("/assignment", controllers.CreateAssignment).Methods("POST")

	r.HandleFunc("/assignment/bulk", controllers.AddMultipleAssignments).Methods("POST")

	r.HandleFunc("/assignment", controllers.GetAssignments).Methods("GET")

	r.HandleFunc("/assignment/{id}", controllers.GetAssignmentByID).Methods("GET")

	r.HandleFunc("/assignment/{id}", controllers.UpdateAssignment).Methods("PUT")

	r.HandleFunc("/assignment/{id}", controllers.DeleteAssignment).Methods("DELETE")

	r.HandleFunc("/submit", controllers.SubmitAssignment).Methods("POST")
}