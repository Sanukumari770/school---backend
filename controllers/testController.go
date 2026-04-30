package controllers
import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func AddTest(w http.ResponseWriter, r *http.Request) {

	var test models.Test
	json.NewDecoder(r.Body).Decode(&test)

	test.ID = primitive.NewObjectID()
	test.CreatedAt = time.Now()

	if test.Marks >= 40 {
		test.Result = "pass"
	} else {
		test.Result = "fail"
	}

	collection := config.DB.Collection("tests")
	collection.InsertOne(context.Background(), test)

	json.NewEncoder(w).Encode(test)
}