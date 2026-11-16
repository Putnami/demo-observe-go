package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.putnami.dev/protocol/features/spectest"
)

func TestQuote(t *testing.T) {
	spectest.Proves(t, "api/quote", "a-quote-includes-tax", "two-five-euro-lines-quote-twelve-euros")
	body := `{"lines":[{"SKU":"a","UnitCents":500,"Quantity":2}]}`
	rec := httptest.NewRecorder()
	quote(rec, httptest.NewRequest(http.MethodPost, "/quote", strings.NewReader(body)))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var got quoteResponse
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.TotalCents != 1200 {
		t.Fatalf("TotalCents = %d, want 1200", got.TotalCents)
	}
	if got.ID == "" {
		t.Fatal("ID is empty")
	}
}

func TestQuoteRejectsNegative(t *testing.T) {
	spectest.Proves(t, "api/quote", "a-negative-line-is-unprocessable", "a-negative-unit-price-answers-422")
	body := `{"lines":[{"SKU":"a","UnitCents":-5,"Quantity":1}]}`
	rec := httptest.NewRecorder()
	quote(rec, httptest.NewRequest(http.MethodPost, "/quote", strings.NewReader(body)))
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want 422", rec.Code)
	}
}

func TestQuoteRejectsGet(t *testing.T) {
	spectest.Proves(t, "api/quote", "only-post-is-accepted", "a-get-answers-405")
	rec := httptest.NewRecorder()
	quote(rec, httptest.NewRequest(http.MethodGet, "/quote", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", rec.Code)
	}
}
