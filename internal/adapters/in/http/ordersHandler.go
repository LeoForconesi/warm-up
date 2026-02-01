package http

import (
	"context"
	"net/http"
	"warm-up/internal/application/ports/in"
	"warm-up/internal/domain"

	"github.com/gin-gonic/gin"
)

// Al crearle una interfaz a OrdersHandlerInterface, se facilita la creación de mocks para pruebas unitarias.
type OrdersHandlerInterface interface {
	FindOrderById(req *gin.Context) domain.Order
	SaveOrder(req *gin.Context) domain.Order
}

type OrdersHandler struct {
	createOrderUC  in.CreateOrderInterface
	getOrderByIdUC in.GetOrderInterface
}

func NewOrdersHandler(createOrderUC in.CreateOrderInterface, getOrderByIdUC in.GetOrderInterface) *OrdersHandler {
	return &OrdersHandler{
		createOrderUC:  createOrderUC,
		getOrderByIdUC: getOrderByIdUC,
	}
}

func (h *OrdersHandler) FindOrderById(req *gin.Context) {
	orderId := req.Param("order_id")

	_ = req.Request.Context()
	_ = context.Background()

	order, err := h.getOrderByIdUC.GetOrderById(orderId)
	if err != nil {
		code, obj := handleError(err)
		req.JSON(code, obj)
		return
	}
	req.JSON(http.StatusOK, order)
}

func handleError(err error) (int, gin.H) {
	switch err.Error() {
	case "order not found":
		return http.StatusNotFound, gin.H{"error": err.Error()}
	case "empty order ID":
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	case "invalid order data":
		return http.StatusBadRequest, gin.H{"error": "amount can't be negative"}
	default:
		return http.StatusInternalServerError, gin.H{"error": "internal server error"}
	}
}

func (h *OrdersHandler) SaveOrder(req *gin.Context) {
	order := domain.Order{}
	if err := req.ShouldBindJSON(&order); err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := h.createOrderUC.CreateOrder(order)
	if err != nil {
		code, obj := handleError(err)
		req.JSON(code, obj)
		return
	}

	req.JSON(http.StatusCreated, newOrder)
}
