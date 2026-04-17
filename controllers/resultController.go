package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Add Result
func AddResult(w http.ResponseWriter, r *http.Request) {

	var data models.Result
	json.NewDecoder(r.Body).Decode(&data)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("results").InsertOne(ctx, data)

	json.NewEncoder(w).Encode(map[string]string{"message": "Result Added"})
}

// Get Results
func GetResults(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := config.DB.Collection("results").Find(ctx, bson.M{})

	var results []models.Result
	cursor.All(ctx, &results)

	json.NewEncoder(w).Encode(results)
}