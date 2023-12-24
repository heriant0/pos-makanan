package categories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type repository struct {
	db *sqlx.DB
}

func newRespository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) GetAll(ctx context.Context) (categoryList []Category, err error) {
	query := `
		SELECT 
			id,
			name,
			description
		FROM 
			categories
	`
	err = r.db.SelectContext(ctx, &categoryList, query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(fmt.Errorf("error repository - GetAll: %w", ErrCategoryNotFound))
			return nil, ErrCategoryNotFound
		}

		return
	}

	if len(categoryList) == 0 {
		log.Error(fmt.Errorf("error repository - GetAll: %w", ErrCategoryNotFound))
		return nil, ErrCategoryNotFound
	}

	return
}
