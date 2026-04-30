

package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"school/config"
	"school/models"
)

// Create Student (Admin)
func AdminCreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)
collection := config.DB.Collection("students")
result, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
json.NewEncoder(w).Encode(result)
}

// Create Class
func CreateClass(w http.ResponseWriter, r *http.Request) {

	var class models.Class
	json.NewDecoder(r.Body).Decode(&class)

	collection := config.DB.Collection("classes")
	res, _ := collection.InsertOne(context.TODO(), class)

	json.NewEncoder(w).Encode(res)
}

// Create Subject
func CreateSubject(w http.ResponseWriter, r *http.Request) {

	var subject models.Subject
	json.NewDecoder(r.Body).Decode(&subject)

	collection := config.DB.Collection("subjects")
	res, _ := collection.InsertOne(context.TODO(), subject)

	json.NewEncoder(w).Encode(res)
}