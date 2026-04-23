package application

import "github.com/allenliao0119/go-auth-service/internal/usecase/api/register"

type UseCase struct{
	User *UserUseCase
}

func NewUseCase(app *Application) *UseCase {
	return &UseCase{
		User: NewUserUseCase(app),
	}
}

type UserUseCase struct{
	Register *register.UseCase
}

func NewUserUseCase(_ *Application) *UserUseCase {
	return &UserUseCase{}
} 