package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"school/config"

	"go.mongodb.org/mongo-driver/bson"
)

// Parent Dashboard
func GetParentDashboard(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id")

	if userID == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	pipeline := []bson.M{

		// find parent by user_id
		{"$match": bson.M{"user_id": userID}},

		// students
		{"$lookup": bson.M{
			"from": "students",
			"localField": "student_ids",
			"foreignField": "_id",
			"as": "students",
		}},

		// attendance
		{"$lookup": bson.M{
			"from": "attendance",
			"localField": "student_ids",
			"foreignField": "student_id",
			"as": "attendance",
		}},

		// marks
		{"$lookup": bson.M{
			"from": "marks",
			"localField": "student_ids",
			"foreignField": "student_id",
			"as": "marks",
		}},

		// fees
		{"$lookup": bson.M{
			"from": "fees",
			"localField": "student_ids",
			"foreignField": "student_id",
			"as": "fees",
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