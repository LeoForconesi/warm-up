package orders

import (
	"time"
	"warm-up/internal/application/ports/out"
	"warm-up/internal/domain"

	"github.com/google/uuid"
)

type CreateOrder struct {
	orderRepo out.OrderRepository
}

func NewCreateOrderUC(orderRepo out.OrderRepository) *CreateOrder {
	return &CreateOrder{
		orderRepo: orderRepo,
	}
}

func (o *CreateOrder) CreateOrder(order domain.Order) (domain.Order, error) {
	if err := validateOrder(order); err != nil {
		return order, err
	}

	order.ID = uuid.New()
	order.CreatedAt = time.Now().Format(time.DateTime)

	err := o.orderRepo.Save(order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func validateOrder(order domain.Order) error {
	if order.Amount <= 0 {
		return ErrInvalidOrderData
	}
	return nil
}
