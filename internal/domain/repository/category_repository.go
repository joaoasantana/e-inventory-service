package repository

import "github.com/joaoasantana/e-inventory-service/internal/domain/entity"

type CategoryRepository interface {
	Create(*entity.Category) error
	FindAll() ([]entity.Category, error)
}
