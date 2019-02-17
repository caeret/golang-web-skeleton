package routing

import (
	"github.com/caeret/golang-web-skeleton/code"
	"github.com/caeret/golang-web-skeleton/request"
	"github.com/caeret/golang-web-skeleton/routing/scope"
	"github.com/caeret/golang-web-skeleton/service"
	routing "github.com/go-ozzo/ozzo-routing"
)

type UserCTL struct {
	userService *service.UserService
}

func (ctl *UserCTL) CreateUser(c *routing.Context) error {
	var req request.CreateUser
	err := c.Read(&req)
	if err != nil {
		return code.NewAPIError("INVALID_DATA").WithDetails(err)
	}
	user, err := ctl.userService.Create(scope.GetRequestScope(c), req)
	if err != nil {
		return err
	}
	return c.Write(user)
}
