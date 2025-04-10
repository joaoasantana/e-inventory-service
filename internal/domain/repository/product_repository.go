package repository

import (
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type ProductRepository interface {
	Create(*entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(uuid.UUID) (*entity.Product, error)
	FindByName(string) (*entity.Product, error)
}
