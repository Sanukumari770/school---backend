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

func AddClass(w http.ResponseWriter, r *http.Request) {

	var class models.Class
	json.NewDecoder(r.Body).Decode(&class)

	class.CreatedAt = time.Now()
	class.UpdatedAt = time.Now()

	res, _ := config.DB.Collection("classes").InsertOne(context.TODO(), class)

	json.NewEncoder(w).Encode(res)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {

	cursor, _ := config.DB.Collection("classes").Find(context.TODO(), bson.M{})

	var classes []models.Class
	cursor.All(context.TODO(), &classes)

	json.NewEncoder(w).Encode(classes)
}