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


// ======================
// CREATE SINGLE FEE
// ======================
func CreateFee(w http.ResponseWriter, r *http.Request) {

	var fee models.Fee

	err := json.NewDecoder(r.Body).Decode(&fee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fee.ID = primitive.NewObjectID()
	fee.DueAmount = fee.TotalAmount - fee.PaidAmount
	fee.Status = "unpaid"
	fee.CreatedAt = time.Now()

	_, err = config.DB.Collection("fees").InsertOne(context.TODO(), fee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Fee Created Successfully")
}


// ======================
// CREATE BULK FEE
// ======================
func CreateBulkFee(w http.ResponseWriter, r *http.Request) {

	var fees []models.Fee

	err := json.NewDecoder(r.Body).Decode(&fees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for i := range fees {
		fees[i].ID = primitive.NewObjectID()
		fees[i].DueAmount = fees[i].TotalAmount - fees[i].PaidAmount
		fees[i].Status = "unpaid"
		fees[i].CreatedAt = time.Now()

		docs = append(docs, fees[i])
	}

	_, err = config.DB.Collection("fees").InsertMany(context.TODO(), docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Bulk Fee Created Successfully")
}


// ======================
// PAY FEE
// ======================
func PayFee(w http.ResponseWriter, r *http.Request) {

	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment.ID = primitive.NewObjectID()
	payment.Status = "success"
	payment.CreatedAt = time.Now()

	res, err := config.DB.Collection("payments").InsertOne(context.TODO(), payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get fee
	var fee models.Fee
	err = config.DB.Collection("fees").
		FindOne(context.TODO(), bson.M{"_id": payment.FeeID}).
		Decode(&fee)
	if err != nil {
		http.Error(w, "Fee not found", http.StatusNotFound)
		return
	}

	newPaid := fee.PaidAmount + payment.Amount
	newDue := fee.TotalAmount - newPaid

	// status logic
	status := "partial"
	if newPaid == 0 {
		status = "unpaid"
	} else if newDue == 0 {
		status = "paid"
	}

	_, err = config.DB.Collection("fees").UpdateOne(
		context.TODO(),
		bson.M{"_id": payment.FeeID},
		bson.M{"$set": bson.M{
			"paid_amount": newPaid,
			"due_amount":  newDue,
			"status":      status,
			"updatedAt":   time.Now(),
		}},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create receipt
	receipt := models.Receipt{
		ID:        primitive.NewObjectID(),
		StudentID: payment.StudentID,
		PaymentID: res.InsertedID.(primitive.ObjectID),
		ReceiptNo: "REC" + time.Now().Format("20060102150405"),
		Amount:    payment.Amount,
		CreatedAt: time.Now(),
	}

	_, err = config.DB.Collection("receipts").InsertOne(context.TODO(), receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Payment Success + Receipt Generated")
}


// ======================
// GET ALL FEES (WITH STUDENT + PAYMENT)
// ======================
func GetAllFees(w http.ResponseWriter, r *http.Request) {

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

		// join payments
		{
			"$lookup": bson.M{
				"from":         "payments",
				"localField":   "_id",
				"foreignField": "fee_id",
				"as":           "payments",
			},
		},
	}

	cursor, err := config.DB.Collection("fees").Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}