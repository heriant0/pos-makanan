package products

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNameIsRequired       = errors.New("name is required")
	ErrStockIsRequired      = errors.New("stock is required")
	ErrPriceIsRequired      = errors.New("price is required")
	ErrCategoryIdIsRequired = errors.New("categoryId is required")
	ErrProductNotFound      = errors.New("product not found")
)

type repository struct {
	db *sqlx.DB
}

func newRespository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, product Product) (id int, err error) {
	query := `
		INSERT INTO  products (
			category_id, name, description, price, stock
		) VALUES (
			:category_id, :name, :description, :price, :stock
		)
		RETURNING id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, product)

	return
}
