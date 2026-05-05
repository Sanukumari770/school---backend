package controllers

import (
	"context"
	"encoding/json"
	"net/http"
"school/config"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo" 
)


// admin dashboard api
func GetDashboard(w http.ResponseWriter, r *http.Request) {

	db := config.DB

	// COUNT DATA
	studentsCount, _ := db.Collection("students").CountDocuments(context.TODO(), bson.M{})
	teachersCount, _ := db.Collection("teachers").CountDocuments(context.TODO(), bson.M{})
	classesCount, _ := db.Collection("classes").CountDocuments(context.TODO(), bson.M{})
	subjectsCount, _ := db.Collection("subjects").CountDocuments(context.TODO(), bson.M{})

	// TOTAL SALARY
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": nil,
				"totalSalary": bson.M{
					"$sum": "$salary",
				},
				"totalBonus": bson.M{
					"$sum": "$bonus",
				},
				"totalDeduction": bson.M{
					"$sum": "$deduction",
				},
			},
		},
	}

	cursor, _ := db.Collection("payroll").Aggregate(context.TODO(), pipeline)

	var salaryData []bson.M
	cursor.All(context.TODO(), &salaryData)

	var totalSalary interface{} = 0
	var totalBonus interface{} = 0
	var totalDeduction interface{} = 0

	if len(salaryData) > 0 {
		totalSalary = salaryData[0]["totalSalary"]
		totalBonus = salaryData[0]["totalBonus"]
		totalDeduction = salaryData[0]["totalDeduction"]
	}

	// FINAL RESPONSE
	response := bson.M{
		"students": studentsCount,
		"teachers": teachersCount,
		"classes":  classesCount,
		"subjects": subjectsCount,
		"salary": bson.M{
		"totalSalary":   totalSalary,
		"totalBonus":    totalBonus,
		"totalDeduction": totalDeduction,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// school full data api students, subjects 
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