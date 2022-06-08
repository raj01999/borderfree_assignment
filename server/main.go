package main

import (
	"log"
	"net/http"
	"server/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.Handle("/home", handlers.AuthenticateToken(http.HandlerFunc(handlers.Home)))
	http.HandleFunc("/refresh", handlers.Refresh)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
