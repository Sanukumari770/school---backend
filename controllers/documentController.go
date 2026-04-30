



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
func UploadDocument(w http.ResponseWriter, r *http.Request) {

	var doc models.Document
	json.NewDecoder(r.Body).Decode(&doc)

	doc.ID = primitive.NewObjectID()
	doc.UploadedAt = time.Now()

	collection := config.DB.Collection("documents")
	collection.InsertOne(context.Background(), doc)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Documents uploaded",
	})
}