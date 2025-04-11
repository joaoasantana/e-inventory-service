package repository

import (
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type SupplierRepository interface {
	Create(*entity.Supplier) error
	FindAll() ([]entity.Supplier, error)
	FindByID(uuid.UUID) (*entity.Supplier, error)
	FindByName(string) (*entity.Supplier, error)
}
