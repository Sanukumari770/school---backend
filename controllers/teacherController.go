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


// =======================
// ✅ ADD TEACHER
// =======================
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


// =======================
// ✅ GET ALL TEACHERS
// =======================
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


// =======================
// ✅ FULL JOIN DATA
// =======================
func GetTeacherFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	teacherID, _ := primitive.ObjectIDFromHex(id)

	pipeline := mongo.Pipeline{

		{{Key: "$match", Value: bson.M{"_id": teacherID}}},

		{{Key: "$lookup", Value: bson.M{
			"from": "subjects",
			"localField": "_id",
			"foreignField": "teacher_id",
			"as": "subjects",
		}}},

		{{Key: "$lookup", Value: bson.M{
			"from": "classes",
			"localField": "classes",
			"foreignField": "_id",
			"as": "classes",
		}}},

		{{Key: "$lookup", Value: bson.M{
			"from": "timetable",
			"localField": "_id",
			"foreignField": "teacher_id",
			"as": "timetable",
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

		// payroll 
		{{Key: "$lookup", Value: bson.M{
			"from": "payroll",
			"localField": "_id",
			"foreignField": "teacher_id",
			"as": "payroll",
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
	}

	cursor, err := config.DB.Collection("teachers").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}


// =======================
// ✅ UPDATE TEACHER
// =======================
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var teacher models.Teacher
	json.NewDecoder(r.Body).Decode(&teacher)

	teacher.UpdatedAt = time.Now()

	_, err := config.DB.Collection("teachers").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": teacher},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Updated")
}


// =======================
// ✅ DELETE (SOFT)
// =======================
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


// =======================
// ✅ DASHBOARD API
// =======================
func GetTeacherDashboard(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	teacherID, _ := primitive.ObjectIDFromHex(id)

	pipeline := mongo.Pipeline{

		{{Key: "$match", Value: bson.M{"_id": teacherID}}},

		{{Key: "$lookup", Value: bson.M{
			"from": "payroll",
			"localField": "_id",
			"foreignField": "teacher_id",
			"as": "payroll",
		}}},

		{{Key: "$lookup", Value: bson.M{
			"from": "classes",
			"localField": "classes",
			"foreignField": "_id",
			"as": "classes",
		}}},
	}

	cursor, _ := config.DB.Collection("teachers").Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}
