package dto

import "github.com/google/uuid"

type ProductRequest struct {
	Name        string  `json:"name"         extensions:"x-order:0" binding:"required"`
	Image       string  `json:"image"        extensions:"x-order:1" binding:"required"`
	Price       float64 `json:"price"        extensions:"x-order:2" binding:"required"`
	Description string  `json:"description"  extensions:"x-order:3" binding:"required"`
}

type ProductResponse struct {
	UUID        uuid.UUID `json:"uuid"          extensions:"x-order:0"`
	Name        string    `json:"name"          extensions:"x-order:1"`
	Image       string    `json:"image"         extensions:"x-order:2"`
	Price       float64   `json:"price"         extensions:"x-order:3"`
	Description string    `json:"description"   extensions:"x-order:4"`
}
