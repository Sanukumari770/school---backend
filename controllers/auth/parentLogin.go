package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"school/config"
	"school/models"
// for security 
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
)

type ParentLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ParentLoginController(w http.ResponseWriter, r *http.Request) {

	var loginData ParentLogin

	json.NewDecoder(r.Body).Decode(&loginData)

	collection := config.DB.Collection("parents")

	var parent models.Parent

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"email": loginData.Email,
		},
	).Decode(&parent)

	if err != nil {

		http.Error(w, "Parent not found", http.StatusUnauthorized)
		return
	}

	inputPassword := strings.TrimSpace(loginData.Password)

	err = bcrypt.CompareHashAndPassword(
		[]byte(parent.Password),
		[]byte(inputPassword),
	)

	if err != nil {

		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// JWT TOKEN

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    parent.ID.Hex(),
		"email": parent.Email,
		"role":  "parent",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {

		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Parent login successful",
		"token":   tokenString,
		"parent":  parent,
	})
}