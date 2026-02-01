package orders

import (
	"fmt"
	"warm-up/internal/application/ports/out"
	"warm-up/internal/domain"
)

type GetOrdersByOrderID struct {
	orderRepo out.OrderRepository
}

func NewGetOrdersByOrderIDUC(orderRepo out.OrderRepository) *GetOrdersByOrderID {
	return &GetOrdersByOrderID{
		orderRepo: orderRepo,
	}
}

func (o *GetOrdersByOrderID) GetOrderById(orderId string) (domain.Order, error) {
	if orderId == "" {
		return domain.Order{}, fmt.Errorf("empty order ID")
	}

	order, err := o.orderRepo.FindById(orderId)
	if err != nil {
		fmt.Println("Error finding order:", err)
		return domain.Order{}, ErrOrderNotFound
	}

	return order, nil
}
