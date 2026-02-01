package orders

import (
	"warm-up/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) Save(order domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepo) FindById(id string) (domain.Order, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func NewMockOrderRepo() *MockOrderRepo {
	return &MockOrderRepo{}
}
