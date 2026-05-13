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

// ==========================
// ADD SINGLE TEACHER
// ==========================

func AddTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	var teacher models.Teacher

	err := json.NewDecoder(r.Body).Decode(&teacher)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teacher.ID = primitive.NewObjectID()

	now := time.Now()

	teacher.CreatedAt = now
	teacher.UpdatedAt = now

	_, err = collection.InsertOne(
		context.Background(),
		teacher,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Teacher Added Successfully",
		"data":    teacher,
	})
}

// ==========================
// ADD MULTIPLE TEACHERS
// ==========================

func AddMultipleTeachers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	var teachers []models.Teacher

	err := json.NewDecoder(r.Body).Decode(&teachers)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for i := range teachers {

		teachers[i].ID = primitive.NewObjectID()

		now := time.Now()

		teachers[i].CreatedAt = now
		teachers[i].UpdatedAt = now

		docs = append(docs, teachers[i])
	}

	result, err := collection.InsertMany(
		context.Background(),
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

// ==========================
// GET ALL TEACHERS
// ==========================

func GetTeachers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	filter := bson.M{
		"deletedAt": bson.M{
			"$exists": false,
		},
	}

	cursor, err := collection.Find(
		context.Background(),
		filter,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var teachers []models.Teacher

	err = cursor.All(context.Background(), &teachers)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if teachers == nil {

		teachers = []models.Teacher{}
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"count":   len(teachers),
		"data":    teachers,
	})
}

// get teacher by id 

func GetTeacherByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var teacher models.Teacher

	err = collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
	).Decode(&teacher)

	if err != nil {

		http.Error(w, "Teacher Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}

// ==========================
// GET FULL TEACHER DATA
// ==========================

func GetTeacherFull(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	id := mux.Vars(r)["id"]


	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	pipeline := bson.A{

		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},

		// SUBJECTS
		bson.M{
			"$lookup": bson.M{
				"from":         "subjects",
				"localField":   "_id",
				"foreignField": "teacher_id",
				"as":           "subjects",
			},
		},

		// CLASSES
		bson.M{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "_id",
				"foreignField": "teacher_id",
				"as":           "classes",
			},
		},

		// ASSIGNMENTS
		bson.M{
			"$lookup": bson.M{
				"from":         "assignments",
				"localField":   "_id",
				"foreignField": "teacher_id",
				"as":           "assignments",
			},
		},

		// PAYROLL
		bson.M{
			"$lookup": bson.M{
				"from":         "payroll",
				"localField":   "_id",
				"foreignField": "teacher_id",
				"as":           "payroll",
			},
		},
	}

	cursor, err := collection.Aggregate(
		context.Background(),
		pipeline,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []bson.M

	err = cursor.All(context.Background(), &result)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}


// updated tecaher 

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updateData bson.M

	err = json.NewDecoder(r.Body).Decode(&updateData)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateData["updated_at"] = time.Now()

	collection := config.DB.Collection("teachers")

    _, err = collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": updateData,
		},
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Teacher Updated Successfully",
	})
}

// deleted teachers 

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("teachers")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	now := time.Now()

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{
				"deleted_at": now,
			},
		},
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Teacher Deleted Successfully",
	})
}