package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"school/utils"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Development mode bypass
		if os.Getenv("APP_ENV") == "development" {

			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 {

			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := splitToken[1]

		token, err := utils.VerifyToken(tokenString)

		if err != nil || !token.Valid {

			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"user_id",
			claims["user_id"],
		)

		ctx = context.WithValue(
			ctx,
			"role",
			claims["role"],
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}