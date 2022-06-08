package main

import (
	"log"
	"net/http"
	h "server/handlers"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/signin", h.Signin).Methods("POST")

	r.HandleFunc("/signup", h.Signup).Methods("POST")

	r.Handle("/addproduct", h.AuthenticateToken(http.HandlerFunc(h.AddProduct))).Methods("POST")

	r.Handle("/deleteproduct", h.AuthenticateToken(http.HandlerFunc(h.DeleteProduct))).Methods("DELETE")

	r.Handle("/getproduct", h.AuthenticateToken(http.HandlerFunc(h.GetProduct))).Methods("GET")

	r.Handle("/updateProduct", h.AuthenticateToken(http.HandlerFunc(h.UpdateProduct))).Methods("PUT")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(r)))
}
