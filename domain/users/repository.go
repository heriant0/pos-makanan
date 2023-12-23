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

func (r repository) Update(ctx context.Context, model User, userId int) (err error) {
	query := `
		UPDATE users
		SET name = :name, date_of_birth = :date_of_birth, phone_number = :phone_number, gender = :gender, address = :address, image_url = :image_url
		WHERE id = :id
`
	model.Id = userId

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
