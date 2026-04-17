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

// Add Fees
func AddFees(w http.ResponseWriter, r *http.Request) {

	var data models.Fees
	json.NewDecoder(r.Body).Decode(&data)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("fees").InsertOne(ctx, data)

	json.NewEncoder(w).Encode(map[string]string{"message": "Fees Added"})
}

// Get Fees
func GetFees(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := config.DB.Collection("fees").Find(ctx, bson.M{})

	var fees []models.Fees
	cursor.All(ctx, &fees)

	json.NewEncoder(w).Encode(fees)
}