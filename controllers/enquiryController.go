// admissiom enquiry 
package controllers

import (
"context"
"encoding/json"
"net/http"
"time"
"school/config"
"school/models"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)
func AddEnquiry(w http.ResponseWriter, r *http.Request) {

	var enquiry models.Enquiry
	json.NewDecoder(r.Body).Decode(&enquiry)

	enquiry.ID = primitive.NewObjectID()
	enquiry.CreatedAt = time.Now()

	collection := config.DB.Collection("enquiries")
	collection.InsertOne(context.Background(), enquiry)

	json.NewEncoder(w).Encode(enquiry)
}

func GetEnquiries(w http.ResponseWriter, r *http.Request) {

	collection := config.DB.Collection("enquiries")

	cursor, _ := collection.Find(context.Background(), bson.M{})
	var enquiries []models.Enquiry
	cursor.All(context.Background(), &enquiries)

	json.NewEncoder(w).Encode(enquiries)

}