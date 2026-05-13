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
// CREATE ENQUIRY
// =============================

func CreateEnquiry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var enquiry models.Enquiry

	err := json.NewDecoder(r.Body).Decode(&enquiry)
	if err != nil {

		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	enquiry.ID = primitive.NewObjectID()

	enquiry.Status = "new"

	enquiry.CreatedAt = time.Now()
	enquiry.UpdatedAt = time.Now()

	collection := config.DB.Collection("enquiries")

	_, err = collection.InsertOne(context.Background(), enquiry)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Enquiry Created Successfully",
		"data":    enquiry,
	})
}


// =============================
// GET ALL ENQUIRIES
// =============================

func GetEnquiries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("enquiries")

	cursor, err := collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	var enquiries []models.Enquiry

	err = cursor.All(context.Background(), &enquiries)
	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"count":   len(enquiries),
		"data":    enquiries,
	})
}


// =============================
// GET ENQUIRY BY ID
// =============================

func GetEnquiryByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	enquiryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var enquiry models.Enquiry

	err = config.DB.
		Collection("enquiries").
		FindOne(
			context.Background(),
			bson.M{"_id": enquiryID},
		).
		Decode(&enquiry)

	if err != nil {

		http.Error(w, "Enquiry Not Found", 404)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"data":    enquiry,
	})
}


// =============================
// UPDATE ENQUIRY STATUS
// =============================

func UpdateEnquiryStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	enquiryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	var data struct {

		Status string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, "Invalid Input", 400)
		return
	}

	_, err = config.DB.
		Collection("enquiries").
		UpdateOne(
			context.Background(),
			bson.M{"_id": enquiryID},
			bson.M{
				"$set": bson.M{
					"status":    data.Status,
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
		"message": "Enquiry Status Updated",
	})
}


// =============================
// DELETE ENQUIRY
// =============================

func DeleteEnquiry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	enquiryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		http.Error(w, "Invalid ID", 400)
		return
	}

	_, err = config.DB.
		Collection("enquiries").
		DeleteOne(
			context.Background(),
			bson.M{"_id": enquiryID},
		)

	if err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Enquiry Deleted Successfully",
	})
}