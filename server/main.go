package main

import (
	"encoding/json"
	"log"
	"net/http"
	h "server/handlers"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"Name":    "Sarafraj Mallick",
		"Email":   "sarafraj01999@gmail.com",
		"Phone":   "7318806855",
		"Detail":  "Server made using GoLang",
		"Message": "Server is running",
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/signin", h.Signin).Methods("POST")

	r.HandleFunc("/signup", h.Signup).Methods("POST")

	r.Handle("/addproduct", h.AuthenticateToken(http.HandlerFunc(h.AddProduct))).Methods("POST")

	r.Handle("/deleteproduct", h.AuthenticateToken(http.HandlerFunc(h.DeleteProduct))).Methods("DELETE")

	r.Handle("/getproduct", h.AuthenticateToken(http.HandlerFunc(h.GetProduct))).Methods("GET")

	r.Handle("/updateProduct", h.AuthenticateToken(http.HandlerFunc(h.UpdateProduct))).Methods("PUT")

	r.HandleFunc("/", home)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(r)))

	log.Fatal(http.ListenAndServe(":"+"5000", handlers.CORS(headers, methods, origins)(r)))
}
