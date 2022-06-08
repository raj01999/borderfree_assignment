package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	crud "server/CRUD_operation"
	"server/model"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value(userId)

	var user Claims
	jsonRead, _ := json.Marshal(props)
	json.Unmarshal(jsonRead, &user)

	var newProduct model.ProductField
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	newProduct.UserId = user.Id

	if newProduct.ProductName == "" || newProduct.ProductDetail == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Please provide the productname and productdetail",
		})
		return
	}

	inserted, err := crud.InsertProduct(newProduct)

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
		"message": "product added successfully",
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value(userId)

	var user Claims
	jsonRead, _ := json.Marshal(props)
	json.Unmarshal(jsonRead, &user)

	var newProduct model.ProductField
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	newProduct.UserId = user.Id

	data := crud.FindOneProduct(newProduct.Id)
	var oldProduct model.ProductField
	data.Decode(&oldProduct)

	if oldProduct.ProductName == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "post not found",
		})
		return
	}

	if oldProduct.UserId != newProduct.UserId {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "user in not the woner",
		})
		return
	}

	deleted, err := crud.DeleteProduct(newProduct.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	fmt.Println(deleted)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "product deleted successfully",
	})

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value(userId)

	var user Claims
	jsonRead, _ := json.Marshal(props)
	json.Unmarshal(jsonRead, &user)

	data, err := crud.FindProduct(user.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	products := []model.ProductField{}
	data.All(context.TODO(), &products)

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]model.ProductField{
		"data": products,
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Response) {

}
