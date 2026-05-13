// token remove 

package middleware
import "net/http"
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Tempo bypass 
		next.ServeHTTP(w, r) // remove token and test api without this 
	})
}