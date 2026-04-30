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

//  Create Student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	res, err := config.DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}

//  Get All Students
func GetStudents(w http.ResponseWriter, r *http.Request) {

	cursor, _ := config.DB.Collection("students").Find(context.TODO(), bson.M{})

	var students []bson.M
	cursor.All(context.TODO(), &students)

	json.NewEncoder(w).Encode(students)
}

// Get Full Student (JOIN)
func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	pipeline := bson.A{

		// student select 
		bson.M{"$match": bson.M{"_id": objID}},

		// class
		bson.M{"$lookup": bson.M{
			"from": "classes",
			"localField": "class_id",
			"foreignField": "_id",
			"as": "class",
		}},

		// subjects
		bson.M{"$lookup": bson.M{
			"from": "subjects",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "subjects",
		}},

		// attendance 
		bson.M{"$lookup": bson.M{
			"from": "attendance",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "attendance",
		}},

		// marks 
		bson.M{"$lookup": bson.M{
			"from": "marks",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "marks",
		}},

		// exam 

		bson.M{"$lookup": bson.M{
			"from": "timetable",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "timetable",
		}},


		// fees
		bson.M{"$lookup": bson.M{
			"from": "fees",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "fees",
		}},

		// books 
		bson.M{"$lookup": bson.M{
			"from": "books",
			"localField": "class_id",
			"foreignField": "class",
			"as": "books",
		}},

		//  Parent
		bson.M{"$lookup": bson.M{
			"from": "parents",
			"localField": "parent_id",
			"foreignField": "_id",
			"as": "parent",
		}},

		//  Assignments
		bson.M{"$lookup": bson.M{
			"from": "assignments",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "assignments",
		}},

		//  Submissions
		bson.M{"$lookup": bson.M{
			"from": "submissions",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "submissions",
		}},
	}

	cursor, _ := config.DB.Collection("students").Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}

//  Update Student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	// take id from URL 
	id := mux.Vars(r)["id"]

	// objectid convert 
	objID, _ := primitive.ObjectIDFromHex(id)

	// take data from body 
	var update bson.M
	json.NewDecoder(r.Body).Decode(&update)

	// add updateAt 
	update["updatedAt"] = time.Now()

	// mongodb update 
	config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)


	// response 
	json.NewEncoder(w).Encode("Updated")
}

//  Delete Student (Soft Delete)
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	now := time.Now()

	config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"deletedAt": now}},
	)

	json.NewEncoder(w).Encode("Deleted")
}