package services

import (
	_"context"
	"gateway/client"
	"gateway/models"
	"gateway/pg" 

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Login(ctx *gin.Context, username, password string) (string, error)
	Register(ctx *gin.Context, req *models.RegisterRequest) (string, error)
}

type userService struct{}

func NewUserService() UserServiceInterface {
	return &userService{}
}

var UserService UserServiceInterface = NewUserService()

func (s *userService) Login(ctx *gin.Context, username, password string) (string, error) {
	resp, err := client.Clients.UserService.Login(ctx.Request.Context(), &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func (s *userService) Register(ctx *gin.Context, req *models.RegisterRequest) (string, error) {
	resp, err := client.Clients.UserService.Register(ctx.Request.Context(), &pb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		RoleId:   int32(req.RoleID),
	})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}