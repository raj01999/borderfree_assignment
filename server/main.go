package main

import (
	"log"
	"net/http"
	"server/handlers"
)

func main() {

	http.HandleFunc("/signin", handlers.Signin)

	http.HandleFunc("/signup", handlers.Signup)

	http.Handle("/addproduct", handlers.AuthenticateToken(http.HandlerFunc(handlers.AddProduct)))

	http.Handle("/deleteproduct", handlers.AuthenticateToken(http.HandlerFunc(handlers.DeleteProduct)))

	http.Handle("/getproduct", handlers.AuthenticateToken(http.HandlerFunc(handlers.GetProduct)))

	http.Handle("/updateProduct", handlers.AuthenticateToken(http.HandlerFunc(handlers.UpdateProduct)))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
