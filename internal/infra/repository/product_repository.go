package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type ProductRepository struct {
	dbConn *sqlx.DB
}

func NewProductRepository(dbConn *sqlx.DB) *ProductRepository {
	return &ProductRepository{dbConn}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	query := `INSERT INTO products (id, name, image, price, description)
		 	  VALUES (:id, :name, :image, :price, :description)`

	result, err := r.dbConn.NamedExec(query, product)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("failed to create product")
	}

	return nil
}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	query := `SELECT id, name, image, price, description
			  FROM products`

	var products []entity.Product

	if err := r.dbConn.Select(&products, query); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id uuid.UUID) (*entity.Product, error) {
	query := `SELECT id, name, image, price, description
			  FROM products
			  WHERE id = $1`

	var product entity.Product

	if err := r.dbConn.Get(&product, query, id); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindByName(name string) (*entity.Product, error) {
	query := `SELECT id, name, image, price, description
			  FROM products
			  WHERE name = $1`

	var product entity.Product

	if err := r.dbConn.Get(&product, query, name); err != nil {
		return nil, err
	}

	return &product, nil
}
