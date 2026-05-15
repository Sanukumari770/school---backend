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


// CREATE ASSIGNMENT

func CreateAssignment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var assignment models.Assignment

	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	assignment.ID = primitive.NewObjectID()

	assignment.CreatedAt = time.Now()

	assignment.UpdatedAt = time.Now()

	// DEFAULT VALUE
	assignment.SubmittedCount = 0

	_, err = config.DB.Collection("assignments").InsertOne(
		context.TODO(),
		assignment,
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Assignment Created Successfully",
	})
}


// ADD MULTIPLE ASSIGNMENTS

func AddMultipleAssignments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var assignments []models.Assignment

	err := json.NewDecoder(r.Body).Decode(&assignments)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	var docs []interface{}

	for i := range assignments {

		assignments[i].ID = primitive.NewObjectID()

		assignments[i].CreatedAt = time.Now()

		assignments[i].UpdatedAt = time.Now()

		assignments[i].SubmittedCount = 0

		docs = append(docs, assignments[i])
	}

	_, err = config.DB.Collection("assignments").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Multiple Assignments Added Successfully",
		"total_assignments_added": len(assignments),
	})
}


// GET ALL ASSIGNMENTS

func GetAssignments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.DB.Collection("assignments").Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(context.TODO())

	var assignments []models.Assignment

	cursor.All(context.TODO(), &assignments)

	if assignments == nil {

		assignments = []models.Assignment{}
	}

	json.NewEncoder(w).Encode(assignments)
}


// GET SINGLE ASSIGNMENT

func GetAssignmentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var assignment models.Assignment

	err = config.DB.Collection("assignments").FindOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	).Decode(&assignment)

	if err != nil {

		http.Error(w, "Assignment not found", 404)
		return
	}

	json.NewEncoder(w).Encode(assignment)
}


// UPDATE ASSIGNMENT

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var assignment models.Assignment

	err = json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	assignment.UpdatedAt = time.Now()

	_, err = config.DB.Collection("assignments").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{

				"title": assignment.Title,

				"subject": assignment.Subject,

				"className": assignment.ClassName,

				"teacherName": assignment.TeacherName,

				"dueDate": assignment.DueDate,

				"totalStudents": assignment.TotalStudents,

				"status": assignment.Status,

				"updatedAt": assignment.UpdatedAt,
			},
		},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Assignment Updated Successfully",
	})
}


// DELETE ASSIGNMENT

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	_, err = config.DB.Collection("assignments").DeleteOne(
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
		"message": "Assignment Deleted Successfully",
	})
}


// SUBMIT ASSIGNMENT

func SubmitAssignment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var submission models.Submission

	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	submission.ID = primitive.NewObjectID()

	submission.SubmittedAt = time.Now()

	_, err = config.DB.Collection("submissions").InsertOne(
		context.TODO(),
		submission,
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	// INCREASE SUBMITTED COUNT

	_, err = config.DB.Collection("assignments").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": submission.AssignmentID,
		},
		bson.M{
			"$inc": bson.M{
				"submittedCount": 1,
			},
		},
	)

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Assignment Submitted Successfully",
	})
}