package orders

import (
	"errors"
	"testing"
	"warm-up/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ordersUseCases(t *testing.T) {
	orderRepo := NewMockOrderRepo()

	t.Run("Get an order by id ok", func(t *testing.T) {
		orderRepo.On("FindById", mock.Anything).Return(
			domain.Order{
				ID:     uuid.MustParse("03a51871-7033-4f9b-a5b1-c8e93349c1bb"),
				Amount: 100.0,
			}, nil).Once()

		getOrderByIdImpl := NewGetOrdersByOrderIDUC(orderRepo)
		res, err := getOrderByIdImpl.GetOrderById("03a51871-7033-4f9b-a5b1-c8e93349c1bb")

		assert.NoError(t, err)
		assert.Equal(t, 100.0, res.Amount)
		assert.Equal(t, "03a51871-7033-4f9b-a5b1-c8e93349c1bb", res.ID.String())
	})

	t.Run("Get an order by id - not found", func(t *testing.T) {
		orderRepo.On("FindById", mock.Anything).Return(
			domain.Order{}, errors.New("order not found")).Once()

		getOrderByIdImpl := NewGetOrdersByOrderIDUC(orderRepo)
		_, err := getOrderByIdImpl.GetOrderById("9999")

		assert.Error(t, err)
		assert.EqualError(t, err, "order not found")
	})

	t.Run("Empty order id", func(t *testing.T) {
		getOrdersByUserID := NewGetOrdersByOrderIDUC(orderRepo)
		_, err := getOrdersByUserID.GetOrderById("")

		assert.Error(t, err)
		assert.EqualError(t, err, "empty order ID")
	})
}
