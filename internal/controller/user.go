package controller

import (
	"github.com/allenliao0119/go-auth-service/internal/usecase/api/register"
	"github.com/gin-gonic/gin"
)

type UserController struct {}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Register(usecase *register.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}