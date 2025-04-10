package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type CategoryRepository struct {
	dbConn *sqlx.DB
}

func NewCategoryRepository(dbConn *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{dbConn}
}

func (r *CategoryRepository) Create(category *entity.Category) error {
	query := `INSERT INTO categories (id, name, description)
		 	  VALUES (:id, :name, :description)`

	result, err := r.dbConn.NamedExec(query, category)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("failed to create category")
	}

	return nil
}

func (r *CategoryRepository) FindAll() ([]entity.Category, error) {
	query := `SELECT id, name, description
			  FROM categories`

	var categories []entity.Category

	if err := r.dbConn.Select(&categories, query); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(id uuid.UUID) (*entity.Category, error) {
	query := `SELECT id, name, description
			  FROM categories
			  WHERE id = $1`

	var category entity.Category

	if err := r.dbConn.Get(&category, query, id); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) FindByName(name string) (*entity.Category, error) {
	query := `SELECT id, name, description
			  FROM categories
			  WHERE name = $1`

	var category entity.Category

	if err := r.dbConn.Get(&category, query, name); err != nil {
		return nil, err
	}

	return &category, nil
}
