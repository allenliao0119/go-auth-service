package register

import (
	"context"
	"errors"
)

type Input struct {
	email string
	password string
}

type Output struct {
	AccessToken string
}

type repository interface {
	EmailExists(email string) (bool, error)
	CreateUser(email, hashedPassword string) (int, error)
}

type password interface{
	Hash(password string) (string, error)
}

type token interface{
	GenerateAccessToken(userID int) (string, error)
}

type UseCase struct {
	repository repository
	password password
	token token
}

func NewUseCase(repository repository, paspassword password, token token) *UseCase {
	return &UseCase{
		repository: repository,
		password: paspassword,
		token: token,
	}
}

func (u *UseCase) Execute(ctx context.Context, input Input) (*Output, error) {
	exists, err := u.repository.EmailExists(input.email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := u.password.Hash(input.password)
	if err != nil {
		return nil, err
	}

	userID, err := u.repository.CreateUser(input.email, hashedPassword)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.token.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	return &Output{
		AccessToken: accessToken,
	}, nil
}