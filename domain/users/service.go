package users

import "context"

type UserRepository interface {
	Update(ctx context.Context, model User, userId int) (err error)
}

type service struct {
	repository UserRepository
}

func newService(repo UserRepository) service {
	return service{
		repository: repo,
	}
}

// func (s service) update(ctx context.Context, req UserRequest) (err error) {

// }
