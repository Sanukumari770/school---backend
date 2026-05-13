package middleware

import (
	"net/http"
)

//  Generic Role Middleware
func Authorize(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			role := r.Context().Value("role")

			// check role exists
			if role == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// match role
			for _, allowed := range allowedRoles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Access denied", http.StatusForbidden)
		})
	}
}