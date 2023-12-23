package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/heriant0/pos-makanan/domain/users"
	"github.com/heriant0/pos-makanan/utility"
)

type RepositoryInterface interface {
	Register(ctx context.Context, auth Auth) (id int, err error)
	GetEmail(ctx context.Context, email string) (string, error)
	GetByEmail(ctx context.Context, email string) (auth Auth, err error)
	Create(ctx context.Context, model users.User) (err error)
	Update(ctx context.Context, id int) (err error)
}

type service struct {
	repository RepositoryInterface
	// userRepo   UserRepository
}

func newService(repo RepositoryInterface) service {
	return service{
		repository: repo,
	}
}

func (s service) register(ctx context.Context, req AuthRequest) (err error) {
	auth, err := requestBody(req)

	if err != nil {
		return
	}

	// check email
	email, _ := s.repository.GetEmail(ctx, auth.Email)
	if email != "" {
		return DuplicateEntry
	}
	// encrypt password
	hashedPassword := utility.HashPassword(auth.Password)
	payload := Auth{
		Email:    auth.Email,
		Password: hashedPassword,
	}

	id, err := s.repository.Register(ctx, payload)
	if err != nil {
		return err
	}

	auth.Id = id
	// create use here
	// userPayload := users.User{
	// 	Id: id,
	// }
	// err = s.repository.Create(ctx, userPayload)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s service) login(ctx context.Context, req AuthRequest) (payload AuthToken, err error) {
	auth, err := requestBody(req)

	if err != nil {
		return
	}

	existingUser, err := s.repository.GetByEmail(ctx, auth.Email)
	if err != nil {
		return AuthToken{}, err
	}

	isVerified := utility.VerifyPassword(req.Password, existingUser.Password)
	if !isVerified {
		return AuthToken{}, errors.New("password verification failed")
	}

	// generate access token
	token, err := utility.GenerateToken(existingUser.Email)
	fmt.Println("🚀 ~ file: service.go ~ line 88 ~ func ~ token : ", token)
	if err != nil {
		return AuthToken{}, err
	}
	payload = AuthToken{
		Access_token: token,
		Role:         existingUser.Role,
	}
	// return c.Status(http.StatusOK).JSON(fiber.Map{
	// 	"token": token,
	// })
	// set token to redis

	return payload, nil
}

func (s service) update(ctx context.Context, id int) (err error) {

	err = s.repository.Update(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
