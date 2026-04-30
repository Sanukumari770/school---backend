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
func AllocateSeat(w http.ResponseWriter, r *http.Request) {

	var seat models.Seat
	json.NewDecoder(r.Body).Decode(&seat)

	seat.ID = primitive.NewObjectID()
	seat.Status = "allocated"
	seat.CreatedAt = time.Now()

	collection := config.DB.Collection("seats")
	collection.InsertOne(context.Background(), seat)

	json.NewEncoder(w).Encode(seat)
}