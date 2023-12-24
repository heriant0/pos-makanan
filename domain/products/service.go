package products

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type RepostoryInterface interface {
	Create(ctx context.Context, product Product, userId int) (id int, err error)
}

type service struct {
	repository RepostoryInterface
}

func newService(repo RepostoryInterface) service {
	return service{
		repository: repo,
	}
}

func (s service) create(ctx context.Context, req ProductRequest, userId int) (err error) {
	product, err := requestBody(req)

	if err != nil {
		log.Error(fmt.Errorf("error service - create: %w", err))

		return
	}

	id, err := s.repository.Create(ctx, product, userId)
	if err != nil {
		log.Error(fmt.Errorf("error service - create: %w", err))
		return err
	}

	product.Id = id

	return nil
}
