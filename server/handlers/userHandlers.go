package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	crud "server/CRUD_operation"
	"server/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string             `json:"username"`
	Id       primitive.ObjectID `json:"_id"`
	jwt.RegisteredClaims
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var newUser model.UserField
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Please Provide The Username And Password",
		})
		return
	}

	result := crud.FindUser(newUser.Username)
	var oldUser model.UserField
	result.Decode(&oldUser)

	if oldUser.Username == newUser.Username {
		w.WriteHeader(http.StatusConflict)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "User Alredy Exist",
		})
		return
	}

	inserted, err := crud.InsertUser(newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	fmt.Println(inserted)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "User Registerd Successfull",
	})

}

func Signin(w http.ResponseWriter, r *http.Request) {
	var newUser model.UserField
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Please Provide The Username And Password",
		})
		return
	}

	result := crud.FindUser(newUser.Username)
	var oldUser model.UserField
	result.Decode(&oldUser)

	if newUser.Username != oldUser.Username {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "User Not Registered",
		})
		return
	}

	if newUser.Password != oldUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Password Is Wrong",
		})
		return
	}

	issueTime := time.Now()
	expirationTime := time.Now().Add(time.Minute * 60)

	claim := &Claims{
		Username: oldUser.Username,
		Id:       oldUser.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issueTime),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "success",
		"message":  "Successfully Login",
		"username": oldUser.Username,
		"token":    tokenString,
	})

}
