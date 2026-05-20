package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"
"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// ADD SINGLE STUDENT

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// HASH PASSWORD
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(student.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	student.Password = string(hashedPassword)

	student.ID = primitive.NewObjectID()

	now := time.Now()

	student.CreatedAt = now
	student.UpdatedAt = now

	_, err = collection.InsertOne(
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




// ADD MULTIPLE STUDENTS

func AddMultipleStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	var students []models.Student

	err := json.NewDecoder(r.Body).Decode(&students)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(students) == 0 {

		http.Error(w, "No students found", http.StatusBadRequest)
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
		"message":      "Multiple Students Added Successfully",
		"inserted_ids": result.InsertedIDs,
	})
}
// GET ALL STUDENTS


func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	filter := bson.M{
		"deleted_at": bson.M{
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


// GET SINGLE STUDENT


func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var student models.Student

	err = collection.FindOne(
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

// GET FULL STUDENT DATA


func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	pipeline := bson.A{

		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},

		// CLASS
		bson.M{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "class_id",
				"foreignField": "_id",
				"as":           "class_data",
			},
		},

		// SUBJECTS
		bson.M{
			"$lookup": bson.M{
				"from":         "subjects",
				"localField":   "class_id",
				"foreignField": "class_id",
				"as":           "subjects",
			},
		},

		// FEES
		bson.M{
			"$lookup": bson.M{
				"from":         "fees",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "fees",
			},
		},

		// TRANSPORT
		bson.M{
			"$lookup": bson.M{
				"from":         "transport",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "transport",
			},
		},

		// ASSIGNMENTS
		bson.M{
			"$lookup": bson.M{
				"from":         "assignments",
				"localField":   "class_id",
				"foreignField": "class_id",
				"as":           "assignments",
			},
		},

		// MARKS
		bson.M{
			"$lookup": bson.M{
				"from":         "marks",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "marks",
			},
		},

		// ATTENDANCE
		bson.M{
			"$lookup": bson.M{
				"from":         "attendance",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "attendance",
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

// UPDATE STUDENT


func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

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
		"message": "Student Updated Successfully",
	})
}


// DELETE STUDENT


func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("students")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
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
		"message": "Student Deleted Successfully",
	})
}