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


// =============================
// CREATE APPLICATION
// =============================

func CreateApplication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var application models.Application

	err := json.NewDecoder(r.Body).Decode(&application)
	if err != nil {

		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	application.ID = primitive.NewObjectID()

	application.ApplicationNo =
		"APP" + time.Now().Format("20060102150405")

	application.Status = "applied"

	application.CreatedAt = time.Now()
	application.UpdatedAt = time.Now()

	collection := config.DB.Collection("applications")

	_, err = collection.InsertOne(context.Background(), application)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Application Created Successfully",
		"data":    application,
	})
}


// =============================
// GET ALL APPLICATIONS
// =============================

func GetApplications(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("applications")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	var applications []models.Application

	err = cursor.All(context.Background(), &applications)
	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"count":   len(applications),
		"data":    applications,
	})
}


// =============================
// GET APPLICATION BY ID
// =============================

func GetApplicationByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	pipeline := []bson.M{

		{
			"$match": bson.M{
				"_id": appID,
			},
		},

		{
			"$lookup": bson.M{
				"from":         "parents",
				"localField":   "parentId",
				"foreignField": "_id",
				"as":           "parent",
			},
		},

		{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "classId",
				"foreignField": "_id",
				"as":           "class",
			},
		},

		{
			"$lookup": bson.M{
				"from":         "students",
				"localField":   "studentName",
				"foreignField": "name",
				"as":           "student",
			},
		},
	}

	cursor, err := config.DB.
		Collection("applications").
		Aggregate(context.Background(), pipeline)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M

	err = cursor.All(context.Background(), &result)
	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"data":    result,
	})
}


// =============================
// UPDATE ENTRANCE RESULT
// =============================

func UpdateEntranceResult(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var data struct {

		EntranceMarks  int    `json:"entranceMarks"`
		EntranceResult string `json:"entranceResult"`
		MeritRank      int    `json:"meritRank"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, "Invalid Input", 400)
		return
	}

	update := bson.M{

		"$set": bson.M{

			"entranceMarks":  data.EntranceMarks,
			"entranceResult": data.EntranceResult,
			"meritRank":      data.MeritRank,

			"status":    "entrance_completed",
			"updatedAt": time.Now(),
		},
	}

	_, err = config.DB.
		Collection("applications").
		UpdateOne(
			context.Background(),
			bson.M{"_id": appID},
			update,
		)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Entrance Result Updated",
	})
}


// =============================
// ALLOCATE SEAT
// =============================

func AllocateSeat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var data struct {

		SeatAllocated bool   `json:"seatAllocated"`
		Section       string `json:"section"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, "Invalid Input", 400)
		return
	}

	update := bson.M{

		"$set": bson.M{

			"seatAllocated": data.SeatAllocated,
			"section":       data.Section,

			"status":    "selected",
			"updatedAt": time.Now(),
		},
	}

	_, err = config.DB.
		Collection("applications").
		UpdateOne(
			context.Background(),
			bson.M{"_id": appID},
			update,
		)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Seat Allocated Successfully",
	})
}


// =============================
// UPDATE FEE STATUS
// =============================

func UpdateFeeStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var data struct {

		FeePaid bool `json:"feePaid"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, "Invalid Input", 400)
		return
	}

	update := bson.M{

		"$set": bson.M{

			"feePaid":  data.FeePaid,
			"updatedAt": time.Now(),
		},
	}

	_, err = config.DB.
		Collection("applications").
		UpdateOne(
			context.Background(),
			bson.M{"_id": appID},
			update,
		)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Fee Updated Successfully",
	})
}


// =============================
// APPROVE APPLICATION
// =============================

func ApproveApplication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var application models.Application

	err = config.DB.
		Collection("applications").
		FindOne(
			context.Background(),
			bson.M{"_id": appID},
		).
		Decode(&application)

	if err != nil {

		http.Error(w, "Application Not Found", 404)
		return
	}

	if application.Status != "selected" {

		http.Error(w, "Seat Not Allocated", 400)
		return
	}

	if !application.FeePaid {

		http.Error(w, "Fee Not Paid", 400)
		return
	}

	student := models.Student{

		ID: primitive.NewObjectID(),

		Name: application.StudentName,

		ClassID: &application.ClassID,

		Email: application.Email,

		Phone: application.Phone,

		CreatedAt: time.Now(),
	}

	res, err := config.DB.
		Collection("students").
		InsertOne(context.Background(), student)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	studentID := res.InsertedID

	_, _ = config.DB.
		Collection("parents").
		UpdateOne(
			context.Background(),
			bson.M{
				"_id": application.ParentID,
			},
			bson.M{
				"$push": bson.M{
					"student_ids": studentID,
				},
			},
		)

	_, _ = config.DB.
		Collection("applications").
		UpdateOne(
			context.Background(),
			bson.M{"_id": appID},
			bson.M{
				"$set": bson.M{
					"status":    "admitted",
					"updatedAt": time.Now(),
				},
			},
		)

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Admission Approved",
		"studentId": studentID,
	})
}


// =============================
// REJECT APPLICATION
// =============================

func RejectApplication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	appID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	_, err = config.DB.
		Collection("applications").
		UpdateOne(
			context.Background(),
			bson.M{"_id": appID},
			bson.M{
				"$set": bson.M{
					"status":    "rejected",
					"updatedAt": time.Now(),
				},
			},
		)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Application Rejected",
	})
}