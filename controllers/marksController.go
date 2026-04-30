package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"school/config"
	"school/models"
)

func AddMarks(w http.ResponseWriter, r *http.Request) {

	var marks models.Marks
	json.NewDecoder(r.Body).Decode(&marks)

	collection := config.DB.Collection("marks")

	res, err := collection.InsertOne(context.TODO(), marks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}