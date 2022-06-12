package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// for passing context {custom type}
type UserKey struct {
	User string
}

// this key use for passing context value from AuthenticateToken middleware to router handlers
var key UserKey = UserKey{"userId"}

// this is an middleware for verifying the token which came from request headers {Authorization}

func AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string = r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "failed",
				"message": "Unauthorized Bad Reques",
			})
			return
		}

		claim := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "failed",
				"message": "Unauthorized",
			})
			return
		}

		// the key use for sending the data after successfully authentication
		ctx := context.WithValue(r.Context(), key, claim)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
