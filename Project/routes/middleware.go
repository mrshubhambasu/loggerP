package routes

import (
	"authservice/utils"
	"net/http"
	"strings"
)

// JwtMiddleware is a middleware for JWT
func JWTMiddleware(next http.Handler) http.Handler {
	tokenValidator := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		_, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(tokenValidator)
}
