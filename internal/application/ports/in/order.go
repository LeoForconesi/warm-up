package in

import "warm-up/internal/domain"

type GetOrderInterface interface {
	GetOrderById(orderId string) (domain.Order, error)
}

type CreateOrderInterface interface {
	CreateOrder(order domain.Order) (domain.Order, error)
}
