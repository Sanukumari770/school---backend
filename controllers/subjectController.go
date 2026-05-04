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

func AddMultipleSubjects(w http.ResponseWriter, r *http.Request) {

	var input []struct {
		Name      string `json:"name"`
		ClassID   string `json:"class_id"`
		TeacherID string `json:"teacher_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var docs []interface{}

	for _, s := range input {

		classID, _ := primitive.ObjectIDFromHex(s.ClassID)
		teacherID, _ := primitive.ObjectIDFromHex(s.TeacherID)

		data := models.Subject{
			Name:      s.Name,
			ClassID:   classID,
			TeacherID: teacherID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		docs = append(docs, data)
	}

	res, _ := config.DB.Collection("subjects").InsertMany(context.TODO(), docs)

	json.NewEncoder(w).Encode(res)
}