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

// Add Class
func AddClass(w http.ResponseWriter, r *http.Request) {

	var data models.Class
	json.NewDecoder(r.Body).Decode(&data)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("classes").InsertOne(ctx, data)

	json.NewEncoder(w).Encode(map[string]string{"message": "Class Added"})
}

// Get Classes
func GetClasses(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := config.DB.Collection("classes").Find(ctx, bson.M{})

	var classes []models.Class
	cursor.All(ctx, &classes)

	json.NewEncoder(w).Encode(classes)
}