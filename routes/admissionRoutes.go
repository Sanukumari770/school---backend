// admission management 

package routes

import (
	"school/controllers"
"github.com/gorilla/mux"
)

func AdmissionRoutes(r *mux.Router) {

	r.HandleFunc("/enquiry", controllers.AddEnquiry).Methods("POST")
	r.HandleFunc("/enquiry", controllers.GetEnquiries).Methods("GET")

	r.HandleFunc("/application", controllers.CreateApplication).Methods("POST")

	r.HandleFunc("/document", controllers.UploadDocument).Methods("POST")

	r.HandleFunc("/test", controllers.AddTest).Methods("POST")

	r.HandleFunc("/merit", controllers.GenerateMerit).Methods("POST")

	r.HandleFunc("/seat", controllers.AllocateSeat).Methods("POST")

	r.HandleFunc("/fee", controllers.PayFee).Methods("POST")

	r.HandleFunc("/admission", controllers.ApproveAdmission).Methods("POST")
}