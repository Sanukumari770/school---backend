package routes

import (
	"school/controllers"

	"github.com/gorilla/mux"
)

func EnquiryRoutes(router *mux.Router) {

	router.HandleFunc(
		"/enquiries",
		controllers.CreateEnquiry,
	).Methods("POST")

	router.HandleFunc(
		"/enquiries",
		controllers.GetEnquiries,
	).Methods("GET")

	router.HandleFunc(
		"/enquiries/{id}",
		controllers.GetEnquiryByID,
	).Methods("GET")

	router.HandleFunc(
		"/enquiries/{id}/status",
		controllers.UpdateEnquiryStatus,
	).Methods("PUT")

	router.HandleFunc(
		"/enquiries/{id}",
		controllers.DeleteEnquiry,
	).Methods("DELETE")
}