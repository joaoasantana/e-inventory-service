package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Stock struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	Quantity  int64
}

func NewStock(id uuid.UUID, productID uuid.UUID, quantity int64) (*Stock, error) {
	stock := &Stock{
		ID:        id,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := stock.validateRules(); err != nil {
		return nil, err
	}

	return stock, nil
}

func (e *Stock) validateRules() error {
	if e.ID == uuid.Nil {
		return errors.New("stock id cannot be null")
	}

	if e.ProductID == uuid.Nil {
		return errors.New("invalid product id")
	}

	if e.Quantity < 0 {
		return errors.New("quantity must be positive")
	}

	return nil
}
