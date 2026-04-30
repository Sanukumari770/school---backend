package main

import (
	"log"
	"net/http"
	"os"

	"school/config"
	"school/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	//  Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	//  Debug env
	log.Println("MONGO_URI:", os.Getenv("MONGO_URI"))

	//  Connect DB
	config.ConnectDB()

	//  Router
	r := mux.NewRouter()

	//  Routes
	routes.SetupRoutes(r)
	routes.AdmissionRoutes(r)

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173", // your frontend
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	//  Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}