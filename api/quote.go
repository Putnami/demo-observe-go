// Package main serves order quotes over HTTP.
package main

import (
	"encoding/json"
	"net/http"

	"github.com/Putnami/demo-observe-go/pricing"
	"github.com/google/uuid"
)

// taxBasisPoints is the tax rate applied to every quote: 20 %.
const taxBasisPoints = 2000

type quoteRequest struct {
	Lines []pricing.Line `json:"lines"`
}

type quoteResponse struct {
	ID         string `json:"id"`
	TotalCents int64  `json:"totalCents"`
}

func quote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req quoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	subtotal, err := pricing.Subtotal(req.Lines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(quoteResponse{
		ID:         uuid.NewString(),
		TotalCents: pricing.WithTax(subtotal, taxBasisPoints),
	})
}
