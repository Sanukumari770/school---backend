package controllers

import (
	"context"
	"encoding/json"
	"net/http"
"school/config"
"go.mongodb.org/mongo-driver/bson"
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