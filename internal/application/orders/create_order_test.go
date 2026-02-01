package orders

import (
	"testing"
	"warm-up/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_createOrder(t *testing.T) {
	orderRepo := NewMockOrderRepo()

	t.Run("should create an order successfully", func(t *testing.T) {
		orderRepo.On("Save", mock.Anything).Return(nil).Once()
		createOrderImpl := NewCreateOrderUC(orderRepo)
		order, err := createOrderImpl.CreateOrder(
			domain.Order{
				Amount: 250.0,
			},
		)

		assert.NoError(t, err)
		assert.NotNil(t, order.ID)
		assert.Equal(t, 250.0, order.Amount)
	})

	t.Run("should return error when order creation fails", func(t *testing.T) {
		orderRepo.On("Save", mock.Anything).Return(assert.AnError).Once()
		createOrderImpl := NewCreateOrderUC(orderRepo)
		_, err := createOrderImpl.CreateOrder(
			domain.Order{
				ID:     uuid.MustParse("03a51871-7033-4f9b-a5b1-c8e93349c1bb"),
				Amount: 300.0,
			},
		)

		assert.Error(t, err)
	})

}
