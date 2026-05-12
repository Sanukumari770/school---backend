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

var teacherCollection = config.DB.Collection("teachers")

// add single teacher

func AddTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

	_, err = teacherCollection.InsertOne(
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
 
 // add multiple teachers

func AddMultipleTeachers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var teachers []models.Teacher

	err := json.NewDecoder(r.Body).Decode(&teachers)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(teachers) == 0 {

		http.Error(w, "No teachers data found", http.StatusBadRequest)
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

	result, err := teacherCollection.InsertMany(
		context.Background(),
		docs,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success":      true,
		"message":      "Multiple Teachers Added Successfully",
		"inserted_ids": result.InsertedIDs,
		"count":        len(result.InsertedIDs),
	})
}


// get all teachers 

func GetTeachers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	filter := bson.M{
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	cursor, err := teacherCollection.Find(
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


// get single teacher 

func GetTeacherByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var teacher models.Teacher

	err = teacherCollection.FindOne(
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

	_, err = teacherCollection.UpdateOne(
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

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	now := time.Now()

	_, err = teacherCollection.UpdateOne(
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