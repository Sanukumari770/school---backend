package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"school/config"
	"school/models"
)

func CreateExam(w http.ResponseWriter, r *http.Request) {

	var exam models.Exam
	json.NewDecoder(r.Body).Decode(&exam)

	collection := config.DB.Collection("exams")

	res, err := collection.InsertOne(context.TODO(), exam)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}