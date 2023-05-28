package main

import (
	"bitcoin-notify/internal/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()
	// Register Routes
	RegisterProductRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/rate", controller.PriceHandler).Methods("GET")
	router.HandleFunc("/api/subscribe", controller.HandleEmail)
	router.HandleFunc("/api/SendEmails", controller.SendEmails)
}
