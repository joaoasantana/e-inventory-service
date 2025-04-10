package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Product struct {
	UUID        uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Image       string    `db:"image"`
	Price       float64   `db:"price"`
	Description string    `db:"description"`
}

func (entity *Product) ValidateRules() error {
	if entity.UUID == uuid.Nil {
		return errors.New("invalid uuid")
	}

	if entity.Name == "" || len(entity.Name) > 255 {
		return errors.New("invalid name")
	}

	if len(entity.Image) > 255 {
		return errors.New("invalid image")
	}

	if entity.Price < 0 {
		return errors.New("price must be positive")
	}

	if entity.Description == "" || len(entity.Description) > 255 {
		return errors.New("invalid description")
	}

	return nil
}
