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

func AddSubject(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name      string `json:"name"`
		ClassID   string `json:"class_id"`
		TeacherID string `json:"teacher_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	classID, _ := primitive.ObjectIDFromHex(input.ClassID)
	teacherID, _ := primitive.ObjectIDFromHex(input.TeacherID)

	data := models.Subject{
		Name:      input.Name,
		ClassID:   classID,
		TeacherID: teacherID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, _ := config.DB.Collection("subjects").InsertOne(context.TODO(), data)

	json.NewEncoder(w).Encode(res)
}