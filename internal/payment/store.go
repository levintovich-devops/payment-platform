package payment

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Store struct {
	mu       sync.RWMutex
	payments map[string]Payment
	seq      uint64
}

func NewStore() *Store {
	return &Store{payments: make(map[string]Payment)}
}

func (s *Store) Create(req PaymentCreateRequest) (Payment, error) {
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
