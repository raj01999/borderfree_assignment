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

var userId UserKey = UserKey{"userId"}

func AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string = r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "failed",
				"message": "unauthorized bad reques",
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
				"message": "unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), userId, claim)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
