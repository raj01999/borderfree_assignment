package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	crud "server/CRUD_operation"
	"server/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// after jwt verification this handler add the product to database with respective user reference

func AddProduct(w http.ResponseWriter, r *http.Request) {

	// the key is use for taking the data from middleware
	props := r.Context().Value(key)

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

	if newProduct.ProductName == "" || newProduct.ProductDetail == "" || newProduct.ProductPrice == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Please Provide The Productname And Productdetail",
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
		"message": "Product Added Successfully",
	})
}

// after token and user varification this handlers send the respective users data to them as and array of object {json}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	// the key is use for taking the data from middleware
	props := r.Context().Value(key)

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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]model.ProductField{
		"data": products,
	})
}

// this handler take a query url parameter id {?id=}, and after verification delete that respective product

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	// the key is use for taking the data from middleware
	props := r.Context().Value(key)

	var user Claims
	jsonRead, _ := json.Marshal(props)
	json.Unmarshal(jsonRead, &user)

	strId := r.URL.Query().Get("id")

	productId, err := primitive.ObjectIDFromHex(strId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := crud.FindOneProduct(productId)
	var oldProduct model.ProductField
	data.Decode(&oldProduct)

	if productId != oldProduct.Id {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Product Not Found",
		})
		return
	}

	if user.Id != oldProduct.UserId {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "User Is Not The Ownler",
		})
		return
	}

	deleted, err := crud.DeleteProduct(productId)

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
		"message": "Product Deleted Successfully",
	})

}

// this handler take a query url parameter id {?id=}, and after verification Update that product respect to req body.

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	// the key is use for taking the data from middleware
	props := r.Context().Value(key)

	var user Claims
	jsonRead, _ := json.Marshal(props)
	json.Unmarshal(jsonRead, &user)

	strId := r.URL.Query().Get("id")

	productId, err := primitive.ObjectIDFromHex(strId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := crud.FindOneProduct(productId)
	var oldProduct model.ProductField
	data.Decode(&oldProduct)

	if productId != oldProduct.Id {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "Product Not Found",
		})
		return
	}

	if user.Id != oldProduct.UserId {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": "User Is Not The Ownler",
		})
		return
	}

	var newProduct model.ProductField
	json.NewDecoder(r.Body).Decode(&newProduct)

	updated, err := crud.UpdateProduct(productId, newProduct)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	fmt.Println(updated)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Product Updated Successfully",
	})

}
