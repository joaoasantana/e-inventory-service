package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
)

type SupplierRepository struct {
	dbConn *sqlx.DB
}

func NewSupplierRepository(dbConn *sqlx.DB) *SupplierRepository {
	return &SupplierRepository{dbConn: dbConn}
}

func (r *SupplierRepository) Create(supplier *entity.Supplier) error {
	query := `INSERT INTO suppliers (id, name, contact)
		 	  VALUES (:id, :name, :contact)`

	result, err := r.dbConn.NamedExec(query, supplier)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("failed to create supplier")
	}

	return nil
}

func (r *SupplierRepository) FindAll() ([]entity.Supplier, error) {
	query := `SELECT id, name, contact
			  FROM suppliers`

	var suppliers []entity.Supplier

	if err := r.dbConn.Select(&suppliers, query); err != nil {
		return nil, err
	}

	return suppliers, nil

}

func (r *SupplierRepository) FindByID(id uuid.UUID) (*entity.Supplier, error) {
	query := `SELECT id, name, contact
			  FROM suppliers
			  WHERE id = $1`

	var supplier entity.Supplier

	if err := r.dbConn.Get(&supplier, query, id); err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *SupplierRepository) FindByName(name string) (*entity.Supplier, error) {
	query := `SELECT id, name, contact
			  FROM suppliers
			  WHERE name = $1`

	var supplier entity.Supplier

	if err := r.dbConn.Get(&supplier, query, name); err != nil {
		return nil, err
	}

	return &supplier, nil
}
