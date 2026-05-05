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
// 1. APPLY ADMISSION
func ApplyAdmission(w http.ResponseWriter, r *http.Request) {

	var admission models.Admission

	err := json.NewDecoder(r.Body).Decode(&admission)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	admission.ID = primitive.NewObjectID()
	admission.ApplicationNo = "APP" + time.Now().Format("20060102150405")
	admission.Status = "pending"
	admission.CreatedAt = time.Now()
	admission.UpdatedAt = time.Now()

	collection := config.DB.Collection("admissions")

	res, err := collection.InsertOne(context.TODO(), admission)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}


// 2. UPDATE ENTRANCE RESULT


func UpdateEntrance(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	admID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var data struct {
		Score  int    `json:"score"`
		Rank   int    `json:"rank"`
		Result string `json:"result"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	update := bson.M{
		"$set": bson.M{
			"entranceScore": data.Score,
			"meritRank":     data.Rank,
			"result":        data.Result,
			"updatedAt":     time.Now(),
		},
	}

	_, err = config.DB.Collection("admissions").UpdateOne(
		context.TODO(),
		bson.M{"_id": admID},
		update,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Entrance Updated")
}


// 3. APPROVE ADMISSION


func ApproveAdmission(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	admID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var admission models.Admission

	err = config.DB.Collection("admissions").
		FindOne(context.TODO(), bson.M{"_id": admID}).
		Decode(&admission)

	if err != nil {
		http.Error(w, "Admission Not Found", 404)
		return
	}

	// create student
	student := models.Student{
		ID:        primitive.NewObjectID(),
		Name:      admission.StudentName,
		ClassID:   &admission.ClassID,
		Email:     admission.Email,
		Phone:     admission.Phone,
		CreatedAt: time.Now(),
	}

	res, err := config.DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	studentID := res.InsertedID

	// link parent
	_, _ = config.DB.Collection("parents").UpdateOne(
		context.TODO(),
		bson.M{"_id": admission.ParentID},
		bson.M{"$push": bson.M{"student_ids": studentID}},
	)

	// transport
	if admission.Transport {
		_, _ = config.DB.Collection("transport").InsertOne(context.TODO(), bson.M{
			"student_id": studentID,
			"bus_no":     "BUS-1",
			"route":      "Default Route",
		})
	}

	//  hostel
	if admission.Hostel {
		_, _ = config.DB.Collection("hostel").InsertOne(context.TODO(), bson.M{
			"student_id": studentID,
			"room_no":    "A-101",
			"block":      "A",
		})
	}

	// update admission status
	_, _ = config.DB.Collection("admissions").UpdateOne(
		context.TODO(),
		bson.M{"_id": admID},
		bson.M{
			"$set": bson.M{
				"status":    "approved",
				"updatedAt": time.Now(),
			},
		},
	)

	json.NewEncoder(w).Encode("Admission Approved & Student Created")
}


// 4. GET FULL ADMISSION (JOIN)


func GetAdmissionFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	admID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	pipeline := []bson.M{

		{"$match": bson.M{"_id": admID}},

		// parent
		{"$lookup": bson.M{
			"from":         "parents",
			"localField":   "parentId",
			"foreignField": "_id",
			"as":           "parent",
		}},

		// class
		{"$lookup": bson.M{
			"from":         "classes",
			"localField":   "classId",
			"foreignField": "_id",
			"as":           "class",
		}},

		// student (after approval)
		{"$lookup": bson.M{
			"from":         "students",
			"localField":   "studentName",
			"foreignField": "name",
			"as":           "student",
		}},
	}

	cursor, err := config.DB.Collection("admissions").
		Aggregate(context.TODO(), pipeline)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}