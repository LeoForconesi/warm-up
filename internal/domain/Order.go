package domain

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID
	Amount    float64
	CreatedAt string
}
