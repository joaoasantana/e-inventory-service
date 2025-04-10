package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/stock/entity"
)

type StockRepositoryDB struct {
	dbConn *sqlx.DB
}

func NewStockRepositoryDB(dbConn *sqlx.DB) *StockRepositoryDB {
	return &StockRepositoryDB{dbConn}
}

func (r *StockRepositoryDB) Create(stock *entity.Stock) error {
	query := `
		INSERT INTO stocks (id, product_id, quantity)
		VALUES (:id, :product_id, :quantity)`

	if _, err := r.dbConn.NamedExec(query, stock); err != nil {
		return err
	}

	return nil
}

func (r *StockRepositoryDB) FindAll() ([]entity.Stock, error) {
	query := `
		SELECT id, product_id, quantity
		FROM stocks`

	var stocks []entity.Stock

	if err := r.dbConn.Select(&stocks, query); err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepositoryDB) FindByID(id uuid.UUID) (*entity.Stock, error) {
	query := `
		SELECT id, product_id, quantity
		FROM stocks
		WHERE id = $1`

	var stock entity.Stock

	if err := r.dbConn.Get(&stock, query, id); err != nil {
		return nil, err
	}

	return &stock, nil
}

func (r *StockRepositoryDB) Update(stock *entity.Stock) error {
	query := `
		UPDATE stocks
		SET quantity = :quantity
		WHERE id = :id`

	if _, err := r.dbConn.NamedExec(query, stock); err != nil {
		return err
	}

	return nil
}

func (r *StockRepositoryDB) Delete(id uuid.UUID) error {
	query := `
		DELETE FROM stocks
		WHERE id = :id`

	if _, err := r.dbConn.NamedExec(query, id); err != nil {
		return err
	}

	return nil
}
