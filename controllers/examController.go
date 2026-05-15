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


// CREATE EXAM

func CreateExam(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var exam models.Exam

	err := json.NewDecoder(r.Body).Decode(&exam)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	exam.ID = primitive.NewObjectID()

	exam.CreatedAt = time.Now()

	exam.UpdatedAt = time.Now()

	_, err = config.DB.Collection("exams").InsertOne(
		context.TODO(),
		exam,
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Exam Created Successfully",
	})
}


// ADD MULTIPLE EXAMS

func AddMultipleExams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var exams []models.Exam

	err := json.NewDecoder(r.Body).Decode(&exams)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	var docs []interface{}

	for i := range exams {

		exams[i].ID = primitive.NewObjectID()

		exams[i].CreatedAt = time.Now()

		exams[i].UpdatedAt = time.Now()

		docs = append(docs, exams[i])
	}

	_, err = config.DB.Collection("exams").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Multiple Exams Added Successfully",
		"total_exams_added": len(exams),
	})
}


// GET ALL EXAMS

func GetExams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.DB.Collection("exams").Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(context.TODO())

	var exams []models.Exam

	cursor.All(context.TODO(), &exams)

	if exams == nil {

		exams = []models.Exam{}
	}

	json.NewEncoder(w).Encode(exams)
}


// GET SINGLE EXAM

func GetExamByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var exam models.Exam

	err = config.DB.Collection("exams").FindOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	).Decode(&exam)

	if err != nil {

		http.Error(w, "Exam not found", 404)
		return
	}

	json.NewEncoder(w).Encode(exam)
}


// UPDATE EXAM

func UpdateExam(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var exam models.Exam

	err = json.NewDecoder(r.Body).Decode(&exam)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	exam.UpdatedAt = time.Now()

	_, err = config.DB.Collection("exams").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{

				"exam_name": exam.ExamName,

				"class_name": exam.ClassName,

				"subject": exam.Subject,

				"exam_date": exam.ExamDate,

				"max_marks": exam.MaxMarks,

				"pass_mark": exam.PassMark,

				"status": exam.Status,

				"updatedAt": exam.UpdatedAt,
			},
		},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Exam Updated Successfully",
	})
}


// DELETE EXAM

func DeleteExam(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	_, err = config.DB.Collection("exams").DeleteOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Exam Deleted Successfully",
	})
}