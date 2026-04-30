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
func GenerateMerit(w http.ResponseWriter, r *http.Request) {

	var merit models.Merit
	json.NewDecoder(r.Body).Decode(&merit)

	merit.ID = primitive.NewObjectID()
	merit.CreatedAt = time.Now()

	collection := config.DB.Collection("merits")
	collection.InsertOne(context.Background(), merit)

	json.NewEncoder(w).Encode(merit)
}