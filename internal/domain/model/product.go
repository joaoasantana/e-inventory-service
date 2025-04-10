package model

import "github.com/google/uuid"

type Product struct {
	UUID        uuid.UUID
	CategoryID  uuid.UUID
	Name        string
	Image       string
	Price       float64
	Description string
}
