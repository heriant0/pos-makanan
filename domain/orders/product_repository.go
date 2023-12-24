package orders

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func newProductRepo(db *sqlx.DB) productRepo {
	return productRepo{db}
}

func (p productRepo) findProductById(ctx context.Context, id int) (product Product, err error) {
	product = Product{
		Name:      "Nasi Goreng Cabe Hijau",
		Price:     21000,
		Stock:     15,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return
}

func (p productRepo) updateProductStockById(ctx context.Context, id int, stock int) (err error) {
	return
}
