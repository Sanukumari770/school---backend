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

	var input []struct {
		Name    string `json:"name"`
		ClassID string `json:"class_id"`
		RollNo  string `json:"roll_no"`
		Class   string `json:"class"`
		Section string `json:"section"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var docs []interface{}

	for _, s := range input {

		classID, err := primitive.ObjectIDFromHex(s.ClassID)
		if err != nil {
			continue // ❌ invalid skip
		}

		data := models.Student{
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

		docs = append(docs, data)
	}

	if len(docs) == 0 {
		http.Error(w, "No valid students to insert", 400)
		return
	}

	res, err := config.DB.Collection("students").InsertMany(context.TODO(), docs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}


// =======================
// GET STUDENTS (SAFE)
// =======================
func GetStudents(w http.ResponseWriter, r *http.Request) {

	cursor, err := config.DB.Collection("students").Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer cursor.Close(context.TODO())

	var students []models.Student

	for cursor.Next(context.TODO()) {
		var s models.Student

		// ❗ skip bad records (important fix)
		if err := cursor.Decode(&s); err != nil {
			continue
		}

		students = append(students, s)
	}

	if students == nil {
		students = []models.Student{}
	}

	json.NewEncoder(w).Encode(students)
}


// =======================
// GET FULL STUDENT (JOIN)
// =======================
func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	pipeline := bson.A{

		bson.M{"$match": bson.M{"_id": objID}},

		bson.M{"$lookup": bson.M{
			"from":         "classes",
			"localField":   "class_id",
			"foreignField": "_id",
			"as":           "class",
		}},

		bson.M{"$lookup": bson.M{
			"from":         "subjects",
			"localField":   "class_id",
			"foreignField": "class_id",
			"as":           "subjects",
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


// =======================
// UPDATE STUDENT
// =======================
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var update bson.M
	json.NewDecoder(r.Body).Decode(&update)

	update["updatedAt"] = time.Now()

	_, err := config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Updated")
}


// =======================
// DELETE STUDENT (SOFT)
// =======================
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	now := time.Now()

	_, err := config.DB.Collection("students").UpdateOne(
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