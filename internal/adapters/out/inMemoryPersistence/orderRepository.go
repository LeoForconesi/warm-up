package inMemoryPersistence

import (
	"errors"
	"sync"
	"warm-up/internal/domain"
)

type InMemoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]domain.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]domain.Order),
	}
}

func (r *InMemoryOrderRepository) Save(order domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID.String()] = order

	return nil
}

func (r *InMemoryOrderRepository) FindById(id string) (domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, exists := r.orders[id]
	if !exists {
		return domain.Order{}, errors.New("element not found") // or return an error if preferred
	}
	return order, nil
}
