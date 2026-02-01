package orders

import "errors"

// Errores del dominio relacionados con órdenes

var (
	// ErrOrderNotFound se devuelve cuando no se encuentra una orden con el ID proporcionado.
	ErrOrderNotFound = errors.New("order not found")

	// ErrInvalidOrderData se devuelve cuando los datos de la orden son inválidos.
	ErrInvalidOrderData = errors.New("invalid order data")

	// ErrOrderCreationFailed se devuelve cuando falla la creación de una orden.
	ErrOrderCreationFailed = errors.New("order creation failed")
)
