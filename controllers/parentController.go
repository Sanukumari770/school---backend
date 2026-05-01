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

// ➤ Create Parent
func CreateParent(w http.ResponseWriter, r *http.Request) {
	var parent models.Parent

	json.NewDecoder(r.Body).Decode(&parent)

	parent.ID = primitive.NewObjectID()
	parent.CreatedAt = time.Now()
	parent.UpdatedAt = time.Now()

	_, err := config.DB.Collection("parents").InsertOne(context.TODO(), parent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(parent)
}

// ➤ Get All Parents
func GetParents(w http.ResponseWriter, r *http.Request) {

	cursor, err := config.DB.Collection("parents").Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var parents []models.Parent
	cursor.All(context.TODO(), &parents)

	json.NewEncoder(w).Encode(parents)
}

// ➤ Get Parent FULL (with Students + data)
func GetParentFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	pipeline := []bson.M{
		{"$match": bson.M{"_id": objID}},

		// join students
		{"$lookup": bson.M{
			"from": "students",
			"localField": "student_ids",
			"foreignField": "_id",
			"as": "students",
		}},

		// join attendance
		{"$lookup": bson.M{
			"from": "attendance",
			"localField": "student_ids",
			"foreignField": "student_id",
			"as": "attendance",
		}},

		// join marks
		{"$lookup": bson.M{
			"from": "marks",
			"localField": "student_ids",
			"foreignField": "student_id",
			"as": "marks",
		}},
	}

	cursor, err := config.DB.Collection("parents").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}

// ➤ Update Parent
func UpdateParent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var parent models.Parent
	json.NewDecoder(r.Body).Decode(&parent)

	parent.UpdatedAt = time.Now()

	update := bson.M{
		"$set": parent,
	}

	_, err := config.DB.Collection("parents").UpdateOne(
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

// ➤ Delete Parent (Soft Delete)
func DeleteParent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"deletedAt": now,
		},
	}

	_, err := config.DB.Collection("parents").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Deleted")
}