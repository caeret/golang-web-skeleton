package service

import (
	"github.com/caeret/golang-web-skeleton/app"
	"github.com/caeret/golang-web-skeleton/code"
	"github.com/caeret/golang-web-skeleton/model"
	"github.com/caeret/golang-web-skeleton/request"
	"golang.org/x/xerrors"
)

type UserService struct{}

func (s *UserService) Create(rs app.RequestScope, request request.CreateUser) (user model.User, err error) {
	exist, err := app.DB.Where("name = ?", request.Name).Exist(new(model.User))
	if err != nil {
		err = xerrors.Errorf("fail to check if user exist: %w", err)
		return
	}
	if exist {
		err = code.NewAPIError(code.UserExist)
		return
	}
	user.Name = request.Name
	user.PasswordHash = []byte(request.Password)
	_, err = app.DB.Insert(&user)
	if err != nil {
		err = xerrors.Errorf("fail to insert user: %w", err)
		return
	}

	return
}
