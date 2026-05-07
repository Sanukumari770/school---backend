package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"school/config"
	"school/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// ======================
// ASSIGN BUS TO STUDENT
// ======================
func AssignTransport(w http.ResponseWriter, r *http.Request) {

	var transport models.Transport

	err := json.NewDecoder(r.Body).Decode(&transport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transport.ID = primitive.NewObjectID()

	_, err = config.DB.Collection("transport").InsertOne(context.TODO(), transport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Transport Assigned Successfully")
}


// ======================
// CREATE BUS
// ======================
func CreateBus(w http.ResponseWriter, r *http.Request) {

	var bus models.Bus

	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bus.ID = primitive.NewObjectID()

	_, err = config.DB.Collection("buses").InsertOne(context.TODO(), bus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Bus Created")
}


// ======================
// GET FULL TRANSPORT DATA (JOIN)
// ======================
func GetTransportDetails(w http.ResponseWriter, r *http.Request) {

	pipeline := []bson.M{

		// join student
		{
			"$lookup": bson.M{
				"from":         "students",
				"localField":   "student_id",
				"foreignField": "_id",
				"as":           "student",
			},
		},

		// join bus (route match)
		{
			"$lookup": bson.M{
				"from":         "buses",
				"localField":   "route",
				"foreignField": "route",
				"as":           "bus",
			},
		},
	}

	cursor, err := config.DB.Collection("transport").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}