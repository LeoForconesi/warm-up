package out

import "warm-up/internal/domain"

type OrderRepository interface {
	Save(order domain.Order) error
	FindById(id string) (domain.Order, error)
}
