package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"

	"go.mongodb.org/mongo-driver/bson"
)

// 📊 Dashboard
func GetDashboard(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	studentCount, _ := config.DB.Collection("students").CountDocuments(ctx, bson.M{})
	teacherCount, _ := config.DB.Collection("teachers").CountDocuments(ctx, bson.M{})
	classCount, _ := config.DB.Collection("classes").CountDocuments(ctx, bson.M{})
	feesCount, _ := config.DB.Collection("fees").CountDocuments(ctx, bson.M{})

	json.NewEncoder(w).Encode(map[string]interface{}{
		"students": studentCount,
		"teachers": teacherCount,
		"classes":  classCount,
		"fees":     feesCount,
	})
}