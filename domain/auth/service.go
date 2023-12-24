package auth

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"github.com/heriant0/pos-makanan/domain/users"
	"github.com/heriant0/pos-makanan/utility"
)

var redisdb *redis.Client

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
		log.Error(fmt.Errorf("error service - register: %w", err))
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
		log.Error(fmt.Errorf("error service - register: %w", err))
		return err
	}

	auth.Id = id
	return nil
}

func (s service) login(ctx context.Context, req AuthRequest) (payload AuthToken, err error) {
	auth, err := requestBody(req)

	if err != nil {
		log.Error(fmt.Errorf("error service - login: %w", err))
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
	tokenObject := AuthTokenPayload{
		Id:    existingUser.Id,
		Email: existingUser.Email,
		Role:  existingUser.Role,
	}

	token, err := utility.GenerateToken(tokenObject)
	if err != nil {
		return AuthToken{}, err
	}

	payload = AuthToken{
		Access_token: token,
		Role:         existingUser.Role,
	}

	// set token to redis
	// key := fmt.Sprintf("%d-%s", existingUser.Id, existingUser.Email)
	// err = utility.SetData(key, token, 10*time.Second)
	// if err != nil {
	// 	return AuthToken{}, err
	// }
	return payload, nil
}

func (s service) update(ctx context.Context, id int) (err error) {

	err = s.repository.Update(ctx, id)
	if err != nil {
		log.Error(fmt.Errorf("error service - update: %w", err))
		return err
	}

	return nil
}
