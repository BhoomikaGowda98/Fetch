package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var receipts = make(map[string]int)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	id := GenerateID()
	points := CalculatePoints(receipt)
	receipts[id] = points

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if points, exists := receipts[id]; exists {
		response := map[string]int{"points": points}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Receipt not found", http.StatusNotFound)
	}
}
