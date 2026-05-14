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

// CREATE BUS

func CreateBus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var bus models.Bus

	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	bus.ID = primitive.NewObjectID()
	bus.CreatedAt = time.Now()

	_, err = config.DB.Collection("buses").InsertOne(
		context.TODO(),
		bus,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Bus Created",
	})
}

// create multiple buses 

func AddMultipleBuses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var buses []models.Bus

	err := json.NewDecoder(r.Body).Decode(&buses)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var docs []interface{}

	for i := range buses {

		buses[i].ID = primitive.NewObjectID()
		buses[i].CreatedAt = time.Now()

		docs = append(docs, buses[i])
	}

	_, err = config.DB.Collection("buses").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Multiple Buses Added",
	})
}

// GET ALL BUSES

func GetBuses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.DB.Collection("buses").Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(context.TODO())

	var buses []models.Bus

	cursor.All(context.TODO(), &buses)

	if buses == nil {
		buses = []models.Bus{}
	}

	json.NewEncoder(w).Encode(buses)
}

// ASSIGN BUS TO students

func AssignTransport(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var transport models.Transport

	err := json.NewDecoder(r.Body).Decode(&transport)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// CHECK BUS EXISTS
	var bus models.Bus

	err = config.DB.Collection("buses").FindOne(
		context.TODO(),
		bson.M{
			"_id": transport.BusID,
		},
	).Decode(&bus)

	if err != nil {
		http.Error(w, "Bus not found", 404)
		return
	}

	// CHECK SEAT AVAILABLE
	if bus.OccupiedSeats >= bus.TotalSeats {
		http.Error(w, "Bus is Full", 400)
		return
	}

	transport.ID = primitive.NewObjectID()
	transport.CreatedAt = time.Now()

	_, err = config.DB.Collection("transport").InsertOne(
		context.TODO(),
		transport,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// INCREASE OCCUPIED SEAT
	_, err = config.DB.Collection("buses").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": transport.BusID,
		},
		bson.M{
			"$inc": bson.M{
				"occupied_seats": 1,
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Transport Assigned Successfully",
	})
}

// FULL TRANSPORT DETAILS

func GetTransportDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pipeline := bson.A{

		// JOIN STUDENT
		bson.M{
			"$lookup": bson.M{
				"from":         "students",
				"localField":   "student_id",
				"foreignField": "_id",
				"as":           "student",
			},
		},


		// JOIN PARENT
bson.M{
	"$lookup": bson.M{
		"from": "parents",
		"localField": "parent_id",
		"foreignField": "_id",
		"as": "parent",
	},
},
		// JOIN BUS
		bson.M{
			"$lookup": bson.M{
				"from":         "buses",
				"localField":   "bus_id",
				"foreignField": "_id",
				"as":           "bus",
			},
		},
	}

	cursor, err := config.DB.Collection("transport").Aggregate(
		context.TODO(),
		pipeline,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M

	cursor.All(context.TODO(), &result)

	if result == nil {
		result = []bson.M{}
	}

	json.NewEncoder(w).Encode(result)
}


// SINGLE BUS DETAILS

func GetBusByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var bus models.Bus

	err = config.DB.Collection("buses").FindOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	).Decode(&bus)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bus)
}