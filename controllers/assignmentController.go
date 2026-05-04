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


// =======================
//  CREATE ASSIGNMENT
// =======================
func CreateAssignment(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title     string `json:"title"`
		SubjectID string `json:"subject_id"`
		TeacherID string `json:"teacher_id"`
		ClassID   string `json:"class_id"`
		DueDate   string `json:"due_date"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	subjectID, _ := primitive.ObjectIDFromHex(input.SubjectID)
	teacherID, _ := primitive.ObjectIDFromHex(input.TeacherID)
	classID, _ := primitive.ObjectIDFromHex(input.ClassID)

	dueDate, _ := time.Parse("2006-01-02", input.DueDate)

	data := models.Assignment{
		Title:     input.Title,
		SubjectID: subjectID,
		TeacherID: teacherID,
		ClassID:   classID,
		DueDate:   dueDate,
		CreatedAt: time.Now(),
	}

	res, _ := config.DB.Collection("assignments").InsertOne(context.TODO(), data)

	json.NewEncoder(w).Encode(res)
}



//  SUBMIT ASSIGNMENT
func SubmitAssignment(w http.ResponseWriter, r *http.Request) {

	var input struct {
		AssignmentID string `json:"assignment_id"`
		StudentID    string `json:"student_id"`
		FileURL      string `json:"file_url"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	assignmentID, _ := primitive.ObjectIDFromHex(input.AssignmentID)
	studentID, _ := primitive.ObjectIDFromHex(input.StudentID)

	data := models.Submission{
		AssignmentID: assignmentID,
		StudentID:    studentID,
		FileURL:      input.FileURL,
		Marks:        0,
		SubmittedAt:  time.Now(),
	}

	res, _ := config.DB.Collection("submissions").InsertOne(context.TODO(), data)

	json.NewEncoder(w).Encode(res)
}