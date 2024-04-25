package main

import (
	"log"
	"net/http"

	"github.com/Nikolay200669/go-ks/internal/app/application"
	iface "github.com/Nikolay200669/go-ks/internal/app/interfaces/http"
)

func main() {
	// Initialize the CalculateService
	calculateService := &application.CalculateService{}

	// Configure the router
	router := iface.SetupRouter(calculateService)

	// Run the server
	log.Fatal(http.ListenAndServe(":8989", router))
}
