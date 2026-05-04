package controllers

import (
	"context"
	"encoding/json"
	"net/http"
"school/config"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo" 
)

// Dashboard
func GetDashboard(w http.ResponseWriter, r *http.Request) {

	db := config.DB

	students, _ := db.Collection("students").CountDocuments(context.TODO(), bson.M{})
	teachers, _ := db.Collection("teachers").CountDocuments(context.TODO(), bson.M{})
	parents, _ := db.Collection("parents").CountDocuments(context.TODO(), bson.M{})
	admissions, _ := db.Collection("admissions").CountDocuments(context.TODO(), bson.M{})

	response := map[string]interface{}{
		"total_students": students,
		"total_teachers": teachers,
		"total_parents": parents,
		"total_admissions": admissions,
	}

	json.NewEncoder(w).Encode(response)
}

func GetSchoolFullData(w http.ResponseWriter, r *http.Request) {

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from": "students",
				"localField": "_id",
				"foreignField": "class_id",
				"as": "students",
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from": "subjects",
				"localField": "_id",
				"foreignField": "class_id",
				"as": "subjects",
			}},
		},
	}

	cursor, err := config.DB.Collection("classes").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}