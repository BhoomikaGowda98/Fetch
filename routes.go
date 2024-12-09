package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// SetupRoutes initializes the application routes.
func SetupRoutes() http.Handler {
	router := chi.NewRouter()
	router.Post("/receipts/process", ProcessReceiptHandler)
	router.Get("/receipts/{id}/points", GetPointsHandler)
	return router
}
