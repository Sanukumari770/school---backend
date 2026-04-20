package main

import (
	"log"
	"net/http"
	"os"

	"school/config"
	"school/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// DB connect
	config.ConnectDB()

	// Router
	r := mux.NewRouter()

	// Routes
	routes.SetupRoutes(r)

	//  CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}