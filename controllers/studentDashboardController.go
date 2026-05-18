package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"school/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStudentDashboard(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := config.DB.Collection("students")

	var student bson.M

	err = collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&student)

	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}