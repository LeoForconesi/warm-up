package dto

// DTOs para las respuestas de la API. Esto es externo al dominio.
type Response struct {
	Data     string `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}
