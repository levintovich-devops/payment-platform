package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/levintovich-devops/payment-platform/internal/payment"
)

func TestCapturePaymentReturnsUpdatedPayment(t *testing.T) {
	store := payment.NewStore()
	created, err := store.Create(payment.PaymentCreateRequest{Reference: "ref-1", Amount: "100", Currency: "USD"})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/payments/"+created.ID+"/capture", nil)
	rec := httptest.NewRecorder()

	NewHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var got payment.Payment
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.ID != created.ID {
		t.Fatalf("expected payment ID %q, got %q", created.ID, got.ID)
	}
	if got.Status != "CAPTURED" {
		t.Fatalf("expected status CAPTURED, got %q", got.Status)
	}
	if got.UpdatedAt == "" {
		t.Fatal("expected updatedAt to be set")
	}
}

func TestCapturePaymentReturnsNotFoundForMissingPayment(t *testing.T) {
	store := payment.NewStore()
	req := httptest.NewRequest(http.MethodPost, "/payments/pay-unknown/capture", nil)
	rec := httptest.NewRecorder()

	NewHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusNotFound, rec.Code, rec.Body.String())
	}

	var errResp payment.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&errResp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if errResp.Code != "payment_not_found" {
		t.Fatalf("expected error code payment_not_found, got %q", errResp.Code)
	}
}
