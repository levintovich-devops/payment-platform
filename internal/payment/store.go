package payment

import (
	"errors"
	"fmt"
	"sort"
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

func (s *Store) Get(id string) (Payment, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	payment, ok := s.payments[id]
	return payment, ok
}

func (s *Store) List(page int, pageSize int) PaymentListResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()

	payments := make([]Payment, 0, len(s.payments))
	for _, payment := range s.payments {
		payments = append(payments, payment)
	}

	sort.Slice(payments, func(i, j int) bool {
		return payments[i].CreatedAt < payments[j].CreatedAt
	})

	totalItems := len(payments)
	if totalItems == 0 {
		return PaymentListResponse{Items: []Payment{}, Page: page, PageSize: pageSize, TotalItems: 0, TotalPages: 0}
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	start := (page - 1) * pageSize
	if start >= totalItems {
		start = totalItems
	}

	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	items := make([]Payment, 0, end-start)
	for _, payment := range payments[start:end] {
		items = append(items, payment)
	}

	return PaymentListResponse{Items: items, Page: page, PageSize: pageSize, TotalItems: totalItems, TotalPages: totalPages}
}
