package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// ➕ Add Student
func AddStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	student.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.DB.Collection("students").InsertOne(ctx, student)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Student Added"})
}

// 📥 Get Students
func GetStudents(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("students").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching data", 500)
		return
	}

	var students []models.Student
	cursor.All(ctx, &students)

	json.NewEncoder(w).Encode(students)
}

// ❌ Delete Student
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("students").DeleteOne(ctx, bson.M{"_id": id})

	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
}