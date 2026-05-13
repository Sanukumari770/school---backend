package routes

import (
	"school/controllers"

	"github.com/gorilla/mux"
)

func ApplicationRoutes(router *mux.Router) {

	router.HandleFunc(
		"/applications",
		controllers.CreateApplication,
	).Methods("POST")

	router.HandleFunc(
		"/applications",
		controllers.GetApplications,
	).Methods("GET")

	router.HandleFunc(
		"/applications/{id}",
		controllers.GetApplicationByID,
	).Methods("GET")

	router.HandleFunc(
		"/applications/{id}/entrance",
		controllers.UpdateEntranceResult,
	).Methods("PUT")

	router.HandleFunc(
		"/applications/{id}/seat",
		controllers.AllocateSeat,
	).Methods("PUT")

	router.HandleFunc(
		"/applications/{id}/fee",
		controllers.UpdateFeeStatus,
	).Methods("PUT")

	router.HandleFunc(
		"/applications/{id}/approve",
		controllers.ApproveApplication,
	).Methods("POST")

	router.HandleFunc(
		"/applications/{id}/reject",
		controllers.RejectApplication,
	).Methods("POST")
}