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

// ➕ Mark Attendance
func MarkAttendance(w http.ResponseWriter, r *http.Request) {

	var data models.Attendance

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.DB.Collection("attendance").InsertOne(ctx, data)
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Attendance Marked",
	})
}

// 📥 Get Attendance
func GetAttendance(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("attendance").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching data", 500)
		return
	}

	var data []models.Attendance
	cursor.All(ctx, &data)

	json.NewEncoder(w).Encode(data)
}