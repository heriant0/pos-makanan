package products

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPriceIsRequired       = errors.New("price is required")
	ErrPriceIsInvalid        = errors.New("price is invalid")
	ErrStockIsRequired       = errors.New("stock is required")
	ErrStockIsInvalid        = errors.New("stock is invalid")
	ErrNameIsRequired        = errors.New("name is required")
	ErrDescriptionIsRequierd = errors.New("description is required")
	ErrImageUrlIsRequierd    = errors.New("image url is required")
	ErrCategoryIdIsRequired  = errors.New("categoryId is required")
	ErrCategoryIdIsNotFound  = errors.New("categoryId is not found")
	ErrInvalidRole           = errors.New("invalid role")
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, product Product, userId int) (id int, err error) {
	query := `
		INSERT INTO  products (
			category_id, name, description, price, stock, image_url
		) VALUES (
			:category_id, :name, :description, :price, :stock, :image_url
		)
		RETURNING id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		log.Error(fmt.Errorf("error repository - Create: %w", err))
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, product)

	return
}
