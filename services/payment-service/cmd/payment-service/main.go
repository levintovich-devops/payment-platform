package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type PaymentCreateRequest struct {
	Reference string `json:"reference"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
}

type Payment struct {
	ID        string `json:"id"`
	Reference string `json:"reference"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type paymentStore struct {
	mu       sync.RWMutex
	payments map[string]Payment
	seq      uint64
}

func newPaymentStore() *paymentStore {
	return &paymentStore{payments: make(map[string]Payment)}
}

func (s *paymentStore) Create(req PaymentCreateRequest) (Payment, error) {
	trimmedReference := strings.TrimSpace(req.Reference)
	trimmedAmount := strings.TrimSpace(req.Amount)
	trimmedCurrency := strings.TrimSpace(req.Currency)

	if trimmedReference == "" || trimmedAmount == "" || trimmedCurrency == "" {
		return Payment{}, errors.New("reference, amount, and currency are required")
	}

	now := time.Now().UTC()
	paymentID := fmt.Sprintf("pay-%d-%d", now.UnixNano(), atomic.AddUint64(&s.seq, 1))
	payment := Payment{
		ID:        paymentID,
		Reference: trimmedReference,
		Amount:    trimmedAmount,
		Currency:  trimmedCurrency,
		Status:    "INITIATED",
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.payments[payment.ID] = payment
	return payment, nil
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func newHandler(store *paymentStore) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req PaymentCreateRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "invalid_request", Message: "request body must be valid JSON"})
				return
			}

			payment, err := store.Create(req)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "invalid_request", Message: err.Error()})
				return
			}

			writeJSON(w, http.StatusCreated, payment)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	return mux
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	mux := newHandler(newPaymentStore())

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server stopped unexpectedly", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info("server started", "addr", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
