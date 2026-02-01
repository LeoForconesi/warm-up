package http

var (
	// ErrInvalidRequest represents an error for invalid requests
	ErrInvalidRequest = "invalid request"

	// ErrOrderNotFound represents an error when an order is not found
	ErrOrderNotFound = "order not found"

	// ErrInternalServer represents a generic internal server error
	ErrInternalServer = "internal server error"
)

//type ErrorResponse struct {
//	Error errors.Error `json:"error"`
//}
