package main

import (
"log"
"net/http"
"os"
"school/config"
"school/routes"
"github.com/gorilla/mux"
"github.com/joho/godotenv"
)
func main() {

	// Load env envrioment setup 
	godotenv.Load()

	// Database connect
	config.ConnectDB()

	// Router convert into MUX 
	r := mux.NewRouter()

	// Routes setup
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}