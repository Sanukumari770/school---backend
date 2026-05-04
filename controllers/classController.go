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
func AddMultipleClasses(w http.ResponseWriter, r *http.Request) {

	var classes []models.Class

	err := json.NewDecoder(r.Body).Decode(&classes)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	for i := range classes {
		classes[i].CreatedAt = time.Now()
		classes[i].UpdatedAt = time.Now()
	}

	var docs []interface{}
	for _, c := range classes {
		docs = append(docs, c)
	}

	res, err := config.DB.Collection("classes").InsertMany(context.TODO(), docs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}


func GetClasses(w http.ResponseWriter, r *http.Request) {

	cursor, _ := config.DB.Collection("classes").Find(context.TODO(), bson.M{})

	var classes []models.Class
	cursor.All(context.TODO(), &classes)

	json.NewEncoder(w).Encode(classes)
}