package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
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

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Salary Added Successfully",
	})
}

	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)

	teacherIDHex := input["teacher_id"].(string)

	teacherID, err := primitive.ObjectIDFromHex(teacherIDHex)
	if err != nil {
		http.Error(w, "Invalid teacher_id", 400)
		return
	}

	data := models.Payroll{
		ID:        primitive.NewObjectID(),
		TeacherID: teacherID,
		Salary:    input["salary"].(float64),
		Bonus:     input["bonus"].(float64),
		Deduction: input["deduction"].(float64),
		Month:     input["month"].(string),
		CreatedAt: time.Now(),
	}

	_, err = config.DB.Collection("payroll").InsertOne(context.TODO(), data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Salary Added")
// get salary
	func GetSalaryByTeacher(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["teacherId"]
		teacherID, _ := primitive.ObjectIDFromHex(id)
	
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
