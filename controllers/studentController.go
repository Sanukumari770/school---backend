// students api testing thriugh this struc 

package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"
// uses mux 
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var studentCollection = config.DB.Collection("students")

// add single students 

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = primitive.NewObjectID() // students id 

	now := time.Now()

	student.CreatedAt = now
	student.UpdatedAt = now

	_, err = studentCollection.InsertOne(
		context.Background(),
		student,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student Added Successfully",
		"data":    student,
	})
}

// add multiple students 

func AddMultipleStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var students []models.Student

	err := json.NewDecoder(r.Body).Decode(&students)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for i := range students {

		students[i].ID = primitive.NewObjectID()

		now := time.Now()

		students[i].CreatedAt = now
		students[i].UpdatedAt = now

		docs = append(docs, students[i])
	}

	result, err := studentCollection.InsertMany(
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
// get students 

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	filter := bson.M{
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	cursor, err := studentCollection.Find(
		context.Background(),
		filter,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var students []models.Student

	err = cursor.All(context.Background(), &students)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if students == nil {

		students = []models.Student{}
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"count":   len(students),
		"data":    students,
	})
}

// get students by id 

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var student models.Student

	err = studentCollection.FindOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
	).Decode(&student)

	if err != nil {

		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}
// updated students 

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updateData bson.M

	err = json.NewDecoder(r.Body).Decode(&updateData)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateData["updated_at"] = time.Now()

	_, err = studentCollection.UpdateOne(
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
		"message": "Student Updated Successfully",
	})
}

// deleted students 

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	now := time.Now()

	_, err = studentCollection.UpdateOne(
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
		"message": "Student Deleted Successfully",
	})
}