package http

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/levintovich-devops/payment-platform/internal/payment"
)

func NewHandler(store *payment.Store) nethttp.Handler {
	mux := nethttp.NewServeMux()
	mux.HandleFunc("/healthz", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		w.WriteHeader(nethttp.StatusOK)
	})
	mux.HandleFunc("/readyz", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		w.WriteHeader(nethttp.StatusOK)
	})
	mux.HandleFunc("/payments", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		switch r.Method {
		case nethttp.MethodPost:
			var req payment.PaymentCreateRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, nethttp.StatusBadRequest, payment.ErrorResponse{Code: "invalid_request", Message: "request body must be valid JSON"})
				return
			}

			createdPayment, err := store.Create(req)
			if err != nil {
				writeJSON(w, nethttp.StatusBadRequest, payment.ErrorResponse{Code: "invalid_request", Message: err.Error()})
				return
			}

			writeJSON(w, nethttp.StatusCreated, createdPayment)
		default:
			w.WriteHeader(nethttp.StatusMethodNotAllowed)
		}
	})
	return mux
}

func writeJSON(w nethttp.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
