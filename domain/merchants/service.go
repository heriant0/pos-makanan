package merchants

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	Create(ctx context.Context, user Merchant, mId int) (id int, err error)
	Update(ctx context.Context, model MerchantRequest, mId int) (err error)
	GetProfile(ctx context.Context, userId int) (user Merchant, err error)
}

type service struct {
	repository UserRepository
}

func newService(repo UserRepository) service {
	return service{
		repository: repo,
	}
}

func (s service) create(ctx context.Context, req MerchantRequest, mId int) (err error) {
	merchant, err := requestBody(req)

	if err != nil {
		return
	}

	_, err = s.repository.Create(ctx, merchant, mId)
	if err != nil {
		log.Error(fmt.Errorf("error service - create: %w", err))

		return err
	}

	return nil
}

func (s service) update(ctx context.Context, req MerchantRequest, userId int) (err error) {
	err = s.repository.Update(ctx, req, userId)

	if err != nil {
		log.Error(fmt.Errorf("error service - update: %w", err))

		return err
	}

	return nil
}

func (s service) getProfile(ctx context.Context, mId int) (merchant Merchant, err error) {

	merchant, err = s.repository.GetProfile(ctx, mId)
	if err != nil {
		log.Error(fmt.Errorf("error service - getProfile: %w", err))

		return merchant, err
	}

	return merchant, nil
}
