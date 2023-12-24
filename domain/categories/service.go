package categories

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type RepostiryInterface interface {
	GetAll(ctx context.Context) (categoryList []Category, err error)
}

type service struct {
	repository RepostiryInterface
}

func newService(repository RepostiryInterface) service {
	return service{
		repository: repository,
	}
}

func (s service) GetAll(ctx context.Context) (categoryList []Category, err error) {
	categoryList, err = s.repository.GetAll(ctx)
	if err != nil {
		log.Error(fmt.Errorf("error service - GetAll: %w", err))

		if err == ErrCategoryNotFound {
			return []Category{}, err
		}
		return nil, err
	}

	return categoryList, nil
}
