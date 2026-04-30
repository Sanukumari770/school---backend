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
func CreateApplication(w http.ResponseWriter, r *http.Request) {

	var app models.Application
	json.NewDecoder(r.Body).Decode(&app)

	app.ID = primitive.NewObjectID()
	app.ApplicationNo = "APP" + time.Now().Format("20060102150405")
	app.Status = "pending"
	app.CreatedAt = time.Now()

	collection := config.DB.Collection("applications")
	collection.InsertOne(context.Background(), app)

	json.NewEncoder(w).Encode(app)
}