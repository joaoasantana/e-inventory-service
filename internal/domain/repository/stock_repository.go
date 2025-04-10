package repository

import (
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/stock/entity"
)

type StockRepository interface {
	Create(*entity.Stock) error
	FindAll() ([]entity.Stock, error)
	FindByID(uuid.UUID) (*entity.Stock, error)
	Update(*entity.Stock) error
	Delete(uuid.UUID) error
}
