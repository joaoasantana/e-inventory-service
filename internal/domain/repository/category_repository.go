package repository

import (
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type CategoryRepository interface {
	Create(*entity.Category) error
	FindAll() ([]entity.Category, error)
	FindByID(uuid.UUID) (*entity.Category, error)
	FindByName(string) (*entity.Category, error)
}
