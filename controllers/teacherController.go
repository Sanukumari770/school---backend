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

// Add Teacher
func AddTeacher(w http.ResponseWriter, r *http.Request) {

	var teacher models.Teacher
	json.NewDecoder(r.Body).Decode(&teacher)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("teachers").InsertOne(ctx, teacher)

	json.NewEncoder(w).Encode(map[string]string{"message": "Teacher Added"})
}

// Get Teachers
func GetTeachers(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := config.DB.Collection("teachers").Find(ctx, bson.M{})

	var teachers []models.Teacher
	cursor.All(ctx, &teachers)

	json.NewEncoder(w).Encode(teachers)
}