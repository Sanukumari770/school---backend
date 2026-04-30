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


// ======================
// CREATE FEE
// ======================
func CreateFee(w http.ResponseWriter, r *http.Request) {

	var fee models.Fee
	json.NewDecoder(r.Body).Decode(&fee)

	fee.ID = primitive.NewObjectID()
	fee.DueAmount = fee.TotalAmount
	fee.Status = "unpaid"
	fee.CreatedAt = time.Now()

	config.DB.Collection("fees").InsertOne(context.TODO(), fee)

	json.NewEncoder(w).Encode("Fee Created")
}


// ======================
// PAY FEE
// ======================
func PayFee(w http.ResponseWriter, r *http.Request) {

	var payment models.Payment
	json.NewDecoder(r.Body).Decode(&payment)

	payment.ID = primitive.NewObjectID()
	payment.Status = "success"
	payment.CreatedAt = time.Now()

	res, _ := config.DB.Collection("payments").InsertOne(context.TODO(), payment)

	// update fee
	var fee models.Fee
	config.DB.Collection("fees").
		FindOne(context.TODO(), bson.M{"_id": payment.FeeID}).
		Decode(&fee)

	newPaid := fee.PaidAmount + payment.Amount
	newDue := fee.TotalAmount - newPaid

	status := "partial"
	if newDue == 0 {
		status = "paid"
	}

	config.DB.Collection("fees").UpdateOne(
		context.TODO(),
		bson.M{"_id": payment.FeeID},
		bson.M{"$set": bson.M{
			"paid_amount": newPaid,
			"due_amount":  newDue,
			"status":      status,
		}},
	)

	// create receipt
	receipt := models.Receipt{
		ID:        primitive.NewObjectID(),
		StudentID: payment.StudentID,
		PaymentID: res.InsertedID.(primitive.ObjectID),
		ReceiptNo: "REC" + time.Now().Format("20060102150405"),
		Amount:    payment.Amount,
		CreatedAt: time.Now(),
	}

	config.DB.Collection("receipts").InsertOne(context.TODO(), receipt)

	json.NewEncoder(w).Encode("Payment Success + Receipt Generated")
}


// ======================
// GET FULL FEE DETAILS (JOIN)
// ======================
func GetFeeDetails(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	studentID, _ := primitive.ObjectIDFromHex(id)

	pipeline := []bson.M{

		{"$match": bson.M{"student_id": studentID}},

		// student
		{"$lookup": bson.M{
			"from":         "students",
			"localField":   "student_id",
			"foreignField": "_id",
			"as":           "student",
		}},

		// payments
		{"$lookup": bson.M{
			"from":         "payments",
			"localField":   "_id",
			"foreignField": "fee_id",
			"as":           "payments",
		}},

		// receipts
		{"$lookup": bson.M{
			"from":         "receipts",
			"localField":   "student_id",
			"foreignField": "student_id",
			"as":           "receipts",
		}},
	}

	cursor, _ := config.DB.Collection("fees").Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}