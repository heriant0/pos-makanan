package users

import "context"

type UserRepository interface {
	Create(ctx context.Context, user User, userId int) (id int, err error)
	Update(ctx context.Context, model UserRequest, userId int) (err error)
	GetProfile(ctx context.Context, userId int) (user User, err error)
}

type service struct {
	repository UserRepository
}

func newService(repo UserRepository) service {
	return service{
		repository: repo,
	}
}

func (s service) create(ctx context.Context, req UserRequest, userId int) (err error) {
	user, err := requestBody(req)

	if err != nil {
		return
	}

	_, err = s.repository.Create(ctx, user, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s service) update(ctx context.Context, user UserRequest, userId int) (err error) {
	err = s.repository.Update(ctx, user, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s service) getProfile(ctx context.Context, userId int) (user User, err error) {

	user, err = s.repository.GetProfile(ctx, userId)
	if err != nil {
		return user, err
	}

	return user, nil
}
