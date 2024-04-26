package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nikolay200669/go-ks/internal/app/application"
	iface "github.com/Nikolay200669/go-ks/internal/app/interfaces/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}
	addr := fmt.Sprintf(":%s", port)

	// Initialize the CalculateService
	calculateService := &application.CalculateService{}

	// Configure the router
	router := iface.SetupRouter(calculateService)

	// Run the server
	log.Fatal(http.ListenAndServe(addr, router))
}
