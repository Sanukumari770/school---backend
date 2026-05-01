package middleware

import (
	"context"
	"net/http"
	"strings"
"github.com/golang-jwt/jwt/v5"
	"school/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		tokenStr := parts[1]

		token, err := utils.VerifyToken(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

	//  FIX HERE
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusUnauthorized)
		return
	}
// store role in request context 
	ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
	ctx = context.WithValue(ctx, "role", claims["role"])

	next.ServeHTTP(w, r.WithContext(ctx))
})
}