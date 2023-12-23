package auth

import (
	"context"
	"errors"

	"github.com/heriant0/pos-makanan/domain/users"
	"github.com/jmoiron/sqlx"
)

var (
	EmailIsRequired     = errors.New("email is required")
	EmailIsInvalid      = errors.New("email is invalid")
	PasswordIsEmpty     = errors.New("password is empty")
	PasswordLength      = errors.New("password length must be greater than equal 6")
	DuplicateEntry      = errors.New("email already used")
	Unauthorized        = errors.New("unauthorized")
	UserAlreadyMerchant = errors.New("user already as a merchant")
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) Register(ctx context.Context, auth Auth) (id int, err error) {
	// set default role
	if auth.Role == "" {
		auth.Role = "user"
	}

	query := `
		INSERT INTO auth (
			email, password, role
		) VALUES (
			:email, :password, :role
		)
		RETURNING id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, auth)

	return
}

func (r repository) GetEmail(ctx context.Context, email string) (string, error) {
	var result string

	query := `
        SELECT email FROM auth WHERE email = :email
    `
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &result, map[string]interface{}{"email": email})
	if err != nil {
		return "", err
	}

	return result, nil
}

func (r repository) GetByEmail(ctx context.Context, email string) (auth Auth, err error) {
	var result Auth
	query := `
        SELECT
			id,
			email,
			password,
			role
		FROM auth
		WHERE email = :email
		LIMIT 1;
    `
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return Auth{}, err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &result, map[string]interface{}{"email": email})
	if err != nil {
		return Auth{}, err
	}

	return result, nil
}

func (r repository) Create(ctx context.Context, model users.User) (err error) {
	query := `
		INSERT INTO users (
			id
		) VALUES (
			:id
		)
	`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(model)

	return
}

func (r repository) Update(ctx context.Context, id int) (err error) {
	query := `
		UPDATE auth
		SET role = 'merchant'
		WHERE id = :id
	`
	_, err = r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	return nil
}
