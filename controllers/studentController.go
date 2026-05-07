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
// ADD MULTIPLE STUDENTS
// =======================

func AddMultipleStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var input []struct {
		Name    string `json:"name"`
		ClassID string `json:"class_id"`
		RollNo  string `json:"roll_no"`
		Class   string `json:"class"`
		Section string `json:"section"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for _, s := range input {

		classID, err := primitive.ObjectIDFromHex(s.ClassID)
		if err != nil {
			continue
		}

		student := models.Student{
			ID:        primitive.NewObjectID(),
			Name:      s.Name,
			ClassID:   &classID,
			RollNo:    s.RollNo,
			Class:     s.Class,
			Section:   s.Section,
			Email:     s.Email,
			Phone:     s.Phone,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		docs = append(docs, student)
	}

	if len(docs) == 0 {
		http.Error(w, "No valid students found", http.StatusBadRequest)
		return
	}

	result, err := config.DB.Collection("students").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success":      true,
		"inserted_ids": result.InsertedIDs,
	})
}


// =======================
// GET STUDENTS
// =======================

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	filter := bson.M{
		"deletedAt": bson.M{
			"$exists": false,
		},
	}

	cursor, err := config.DB.Collection("students").Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(context.TODO())

	var students []bson.M

	if err = cursor.All(context.TODO(), &students); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if students == nil {
		students = []bson.M{}
	}

	json.NewEncoder(w).Encode(students)
}


// =======================
// GET FULL STUDENT
// =======================

func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	pipeline := bson.A{

		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "class_id",
				"foreignField": "_id",
				"as":           "class_data",
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         "subjects",
				"localField":   "class_id",
				"foreignField": "class_id",
				"as":           "subjects",
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         "attendance",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "attendance",
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         "marks",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "marks",
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         "fees",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "fees",
			},
		},
	}

	cursor, err := config.DB.Collection("students").Aggregate(
		context.TODO(),
		pipeline,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(result)
}


// =======================
// UPDATE STUDENT
// =======================

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var update bson.M

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	update["updatedAt"] = time.Now()

	_, err = config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": update,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student updated",
	})
}


// =======================
// DELETE STUDENT
// =======================

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	now := time.Now()

	_, err = config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{
				"deletedAt": now,
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student deleted",
	})
}