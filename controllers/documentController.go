package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"school/config"
	"school/models"
	"time"
)

func UploadDocument(w http.ResponseWriter, r *http.Request) {

	var doc models.Document
	json.NewDecoder(r.Body).Decode(&doc)

	doc.ID = primitive.NewObjectID()
	doc.UploadedAt = time.Now()

	collection := config.DB.Collection("documents")
	collection.InsertOne(context.Background(), doc)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Documents uploaded",
	})
}
