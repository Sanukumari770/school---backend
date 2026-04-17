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

// ➕ Add Assignment
func AddAssignment(w http.ResponseWriter, r *http.Request) {

	var data models.Assignment

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.DB.Collection("assignments").InsertOne(ctx, data)
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Assignment Added",
	})
}

// 📥 Get Assignments
func GetAssignments(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("assignments").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching data", 500)
		return
	}

	var data []models.Assignment
	cursor.All(ctx, &data)

	json.NewEncoder(w).Encode(data)
}