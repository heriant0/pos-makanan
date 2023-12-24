package merchants

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	NameIsRequired     = errors.New("name is required")
	AddressIsRequired  = errors.New("address is required")
	PhoneNumberIsEmpty = errors.New("phone number is empty")
	PhoneNumberLength  = errors.New("phone number length must be greater than equal 10")
	ImageUrlIsRequird  = errors.New("image url is required")
	CityIsRequired     = errors.New("city is required")
	TokenEmpty         = errors.New("please provide jwt token")
	InvalidRole        = errors.New("invalid role")
	ErrorRepository    = errors.New("error repository")
	UnknownError       = errors.New("unknown error")
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, model Merchant, mId int) (id int, err error) {
	query := `
		INSERT INTO merchants (
			id, name, address, phone_number, city, image_url
		) VALUES (
			:id, :name, :address, :phone_number, :city, :image_url
		)
		RETURNING id
	`
	model.Id = mId
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, model)

	return
}

func (r repository) Update(ctx context.Context, model MerchantRequest, mId int) (err error) {
	query := `
		UPDATE merchants
		SET name = :name, 
			address = :address, 
			phone_number = :phone_number, 
			city = :city, 
			image_url = :image_url
		WHERE id = :id
`
	_, err = r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"name":         model.Name,
		"address":      model.Address,
		"phone_number": model.PhoneNumber,
		"city":         model.City,
		"image_url":    model.ImageUrl,
		"id":           mId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetProfile(ctx context.Context, mId int) (merchant Merchant, err error) {
	var result Merchant

	query := `
		SELECT 
			name, address, phone_number, city, image_url
		FROM merchants WHERE id = :id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return Merchant{}, err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &result, map[string]interface{}{"id": mId})
	if err != nil {
		return Merchant{}, err
	}

	return result, nil
}
