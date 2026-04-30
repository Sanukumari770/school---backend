// join teacher to relate with class , subject, assigment, attendance, payroll 

package controllers

import (
"context"
"encoding/json"
"net/http"
"time"
"school/config"
"school/models"
"github.com/gorilla/mux"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)

//  Create Teacher (Admin use)
func AddTeacher(w http.ResponseWriter, r *http.Request) {

	var teacher models.Teacher
	json.NewDecoder(r.Body).Decode(&teacher)

	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()

	res, err := config.DB.Collection("teachers").InsertOne(context.TODO(), teacher)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}

//  Get All Teachers
func GetTeachers(w http.ResponseWriter, r *http.Request) {

	cursor, err := config.DB.Collection("teachers").Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var teachers []models.Teacher
	cursor.All(context.TODO(), &teachers)

	json.NewEncoder(w).Encode(teachers)
}

// get single teacher full join data 
func GetTeacherFull(w http.ResponseWriter, r *http.Request) {

	// from URL to find  teacher ID 
	id := mux.Vars(r)["id"]
	teacherID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	// Aggregation Pipeline (JOIN SYSTEM)
	pipeline := mongo.Pipeline{

		// teacher match
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: teacherID},
		}}},

		// subjects (teacher which subject teaches )
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "subjects"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "subjects"},
		}}},

		//classes (teacher which classes teach)
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "classes"},
			{Key: "localField", Value: "classes"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "classes"},
		}}},

		// timetable for classes 
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "timetable"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "timetable"},
		}}},

		// attendance (teacher marks attendance of student )
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "attendance"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "attendance"},
		}}},

		//exams
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "exams"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "exams"},
		}}},

		//assignments
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "assignments"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "assignments"},
		}}},

		//events
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "events"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "events"},
		}}},

		// payroll
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "payroll"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "teacher_id"},
			{Key: "as", Value: "payroll"},
		}}},
	}

	//  aggregation run 
	cursor, err := config.DB.Collection("teachers").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//  result decode
	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// response
	json.NewEncoder(w).Encode(result)
}

// Update Teacher
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var teacher models.Teacher
	json.NewDecoder(r.Body).Decode(&teacher)

	teacher.UpdatedAt = time.Now()

	update := bson.M{
		"$set": teacher,
	}

	_, err := config.DB.Collection("teachers").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Updated")
}

// Delete Teacher (Soft Delete)
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	now := time.Now()

	_, err := config.DB.Collection("teachers").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"deletedAt": now}},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Deleted")
}