package products

import (
	"time"
)

type Product struct {
	Id          int       `db:"id" json:"id"`
	CategoryId  int       `db:"category_id" json:"categoryId"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Stock       int       `db:"stock" json:"stock"`
	ImageUrl    string    `db:"image_url" json:"image_url"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

func requestBody(req ProductRequest) (product Product, err error) {
	product = Product{
		CategoryId:  req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageUrl:    req.ImageUrl,
	}

	err = product.validate()
	return
}

func (p Product) validate() error {
	if err := p.validateName(); err != nil {
		return err
	} else if err := p.validateStock(); err != nil {
		return err
	} else if err = p.validateCategoryId(); err != nil {
		return err
	} else if err := p.validatePrice(); err != nil {
		return err
	}

	return nil
}

func (p Product) validateName() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}
	return nil
}

func (p Product) validateStock() error {
	if p.Stock == 0 {
		return ErrStockIsRequired
	}

	return nil
}

func (p Product) validatePrice() error {
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	return nil
}

func (p Product) validateCategoryId() error {
	if p.CategoryId == 0 {
		return ErrCategoryIdIsRequired
	}

	return nil
}
