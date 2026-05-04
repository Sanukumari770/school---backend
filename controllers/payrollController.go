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
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// =======================
// ✅ ADD SALARY
// =======================
func AddSalary(w http.ResponseWriter, r *http.Request) {

	var input struct {
		TeacherID string  `json:"teacher_id"`
		Salary    float64 `json:"salary"`
		Bonus     float64 `json:"bonus"`
		Deduction float64 `json:"deduction"`
		Month     string  `json:"month"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}

	teacherID, err := primitive.ObjectIDFromHex(input.TeacherID)
	if err != nil {
		http.Error(w, "Invalid teacher_id", 400)
		return
	}

	data := models.Payroll{
		ID:        primitive.NewObjectID(),
		TeacherID: teacherID,
		Salary:    input.Salary,
		Bonus:     input.Bonus,
		Deduction: input.Deduction,
		Month:     input.Month,
		CreatedAt: time.Now(),
	}

	_, err = config.DB.Collection("payroll").InsertOne(context.TODO(), data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Salary Added")
}


// =======================
// ✅ GET SALARY
// =======================
func GetSalaryByTeacher(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["teacherId"]

	teacherID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	cursor, err := config.DB.Collection("payroll").Find(
		context.TODO(),
		bson.M{"teacher_id": teacherID},
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []models.Payroll
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}