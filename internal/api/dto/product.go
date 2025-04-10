package dto

import "github.com/google/uuid"

type ProductRequest struct {
	CategoryID  uuid.UUID `json:"categoryId"    extensions:"x-order:0" binding:"required"`
	Name        string    `json:"name"          extensions:"x-order:1" binding:"required"`
	Image       string    `json:"image"         extensions:"x-order:2" binding:"required"`
	Price       float64   `json:"price"         extensions:"x-order:3" binding:"required"`
	Description string    `json:"description"   extensions:"x-order:4" binding:"required"`
}

type ProductResponse struct {
	UUID        uuid.UUID `json:"id"            extensions:"x-order:0"`
	Category    uuid.UUID `json:"categoryId"    extensions:"x-order:1"`
	Name        string    `json:"name"          extensions:"x-order:2"`
	Image       string    `json:"image"         extensions:"x-order:3"`
	Price       float64   `json:"price"         extensions:"x-order:4"`
	Description string    `json:"description"   extensions:"x-order:5"`
}

type ProductDetailResponse struct {
	UUID        uuid.UUID        `json:"id"            extensions:"x-order:0"`
	Category    CategoryResponse `json:"category"      extensions:"x-order:1"`
	Name        string           `json:"name"          extensions:"x-order:2"`
	Image       string           `json:"image"         extensions:"x-order:3"`
	Price       float64          `json:"price"         extensions:"x-order:4"`
	Description string           `json:"description"   extensions:"x-order:5"`
}
