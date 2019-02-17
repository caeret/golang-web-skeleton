package routing

import (
	"github.com/caeret/golang-web-skeleton/app"
	"github.com/caeret/golang-web-skeleton/request"
	"github.com/caeret/golang-web-skeleton/service"
	routing "github.com/go-ozzo/ozzo-routing"
	"golang.org/x/xerrors"
)

type UserCTL struct {
	userService *service.UserService
}

func (ctl *UserCTL) CreateUser(c *routing.Context) error {
	var req request.CreateUser
	err := c.Read(&req)
	if err != nil {
		return xerrors.Errorf("fail to read request: %v", err)
	}
	user, err := ctl.userService.Create(app.GetRequestScope(c), req)
	if err != nil {
		return err
	}
	return c.Write(user)
}
