package http

import (
	"encoding/json"
	nethttp "net/http"
	"strings"
  "strconv"

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
	mux.HandleFunc("/", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		switch {
		case r.URL.Path == "/payments" && r.Method == nethttp.MethodGet:
			page := 1
			pageSize := 20

			if pageValue := r.URL.Query().Get("page"); pageValue != "" {
				parsedPage, err := strconv.ParseUint(pageValue, 10, 64)
				if err != nil || parsedPage < 1 {
					writeJSON(w, nethttp.StatusBadRequest, payment.ErrorResponse{Code: "invalid_request", Message: "page must be a positive integer"})
					return
				}
				page = int(parsedPage)
			}

			if pageSizeValue := r.URL.Query().Get("pageSize"); pageSizeValue != "" {
				parsedPageSize, err := strconv.ParseUint(pageSizeValue, 10, 64)
				if err != nil || parsedPageSize < 1 || parsedPageSize > 100 {
					writeJSON(w, nethttp.StatusBadRequest, payment.ErrorResponse{Code: "invalid_request", Message: "pageSize must be between 1 and 100"})
					return
				}
				pageSize = int(parsedPageSize)
			}

			writeJSON(w, nethttp.StatusOK, store.List(page, pageSize))
		case r.URL.Path == "/payments":
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
		case strings.HasPrefix(r.URL.Path, "/payments/"):
			if r.Method != nethttp.MethodGet {
				w.WriteHeader(nethttp.StatusMethodNotAllowed)
				return
			}

			paymentID := strings.TrimPrefix(r.URL.Path, "/payments/")
			if paymentID == "" {
				writeJSON(w, nethttp.StatusNotFound, payment.ErrorResponse{Code: "payment_not_found", Message: "payment not found"})
				return
			}

			foundPayment, ok := store.Get(paymentID)
			if !ok {
				writeJSON(w, nethttp.StatusNotFound, payment.ErrorResponse{Code: "payment_not_found", Message: "payment not found"})
				return
			}

			writeJSON(w, nethttp.StatusOK, foundPayment)
		default:
			w.WriteHeader(nethttp.StatusNotFound)
		}
	})
	return mux
}

func writeJSON(w nethttp.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
