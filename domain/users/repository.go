package users

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	GenderIsRequired      = errors.New("gender is required")
	GenderIsInvalid       = errors.New("gender is invalid")
	PhoneNumberIsEmpty    = errors.New("phone number is empty")
	PhoneNumberLength     = errors.New("phone number length must be greater than equal 10")
	NameIsRequired        = errors.New("name is required")
	AddressIsRequired     = errors.New("address is required")
	DateOfBirthIsRequired = errors.New("date of birth is required")
	DateOfBirthIsInvalid  = errors.New("date of birth is invalid")
	ImageUrlIsRequird     = errors.New("image url is required")
	TokenEmpty            = errors.New("please provide jwt token")
	InvalidRole           = errors.New("invalid role")
	ErrorRepository       = errors.New("error repository")
	UnknownError          = errors.New("unknown error")
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, user User, userId int) (id int, err error) {
	query := `
		INSERT INTO users (
			id, name, date_of_birth, phone_number, gender, address, image_url
		) VALUES (
			:id, :name, :date_of_birth, :phone_number, :gender, :address, :image_url
		)
		RETURNING id
	`
	user.Id = userId
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, user)

	return
}

func (r repository) Update(ctx context.Context, user UserRequest, userId int) (err error) {
	query := `
		UPDATE users
		SET name = :name, 
			date_of_birth = :date_of_birth, 
			phone_number = :phone_number, 
			gender = :gender, 
			address = :address, 
			image_url = :image_url
		WHERE id = :id
`
	_, err = r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"name":          user.Name,
		"date_of_birth": user.DateOfBirth,
		"phone_number":  user.PhoneNumber,
		"gender":        user.Gender,
		"address":       user.Address,
		"image_url":     user.ImageUrl,
		"id":            userId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetProfile(ctx context.Context, userId int) (user User, err error) {
	var result User

	query := `
		SELECT 
			name, date_of_birth, phone_number, gender, address, image_url
		FROM users WHERE id = :id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &result, map[string]interface{}{"id": userId})
	if err != nil {
		return User{}, err
	}

	return result, nil
}
