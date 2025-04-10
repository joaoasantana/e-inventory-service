package dto

import "github.com/google/uuid"

type CategoryRequest struct {
	Name        string `json:"name"         extensions:"x-order:0" binding:"required"`
	Description string `json:"description"  extensions:"x-order:1" binding:"required"`
}

type CategoryResponse struct {
	UUID        uuid.UUID `json:"id"            extensions:"x-order:0"`
	Name        string    `json:"name"          extensions:"x-order:1"`
	Description string    `json:"description"   extensions:"x-order:2"`
}
