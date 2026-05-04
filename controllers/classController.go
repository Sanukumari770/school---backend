package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"
"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

// Add Class
func AddClass(w http.ResponseWriter, r *http.Request) {

	var data models.Class
	json.NewDecoder(r.Body).Decode(&data)

	data.ID = primitive.NewObjectID()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	config.DB.Collection("classes").InsertOne(context.TODO(), data)

	json.NewEncoder(w).Encode(data)
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