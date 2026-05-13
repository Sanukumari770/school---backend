package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal(".env file not found")
	}

	requiredEnvs := []string{
		"PORT",
		"MONGO_URI",
		"JWT_SECRET",
		"DB_NAME",
	}

	for _, env := range requiredEnvs {

		if os.Getenv(env) == "" {

			log.Fatal(env + " is missing in .env")
		}
	}
}