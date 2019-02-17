package routing

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caeret/golang-web-skeleton/routing/scope"

	"github.com/caeret/golang-web-skeleton/code"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/fault"

	"github.com/caeret/golang-web-skeleton/service"

	"github.com/caeret/golang-web-skeleton/app"
	routing "github.com/go-ozzo/ozzo-routing"
)

func Serve(logger app.Logger) {
	userService := &service.UserService{}
	userCtl := &UserCTL{
		userService: userService,
	}

	router := routing.New()
	router.Use(
		prepare(logger),
	)
	router.Post("/users", userCtl.CreateUser)
	logger.Info("serve http server.")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func prepare(logger app.Logger) routing.Handler {
	return func(c *routing.Context) error {
		_ = content.TypeNegotiator(content.JSON)(c)
		now := time.Now()
		c.Response = &access.LogResponseWriter{c.Response, http.StatusOK, 0}
		rs := scope.NewRequestScope(now, logger, c.Request)
		c.Set("Context", rs)
		_ = fault.Recovery(func(format string, a ...interface{}) {
			rs.Crit(fmt.Sprintf(format, a...))
		}, convertError(rs))(c)
		logAccess(c, func(format string, a ...interface{}) {
			rs.Info(fmt.Sprintf(format, a...))
		}, rs.Now())
		return nil
	}
}

func logAccess(c *routing.Context, logFunc access.LogFunc, start time.Time) {
	rw := c.Response.(*access.LogResponseWriter)
	elapsed := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	requestLine := fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto)
	logFunc(`%.3fms %s %d %d`, elapsed, requestLine, rw.Status, rw.BytesWritten)
}

func convertError(logger app.Logger) func(*routing.Context, error) error {
	return func(c *routing.Context, err error) error {
		switch err.(type) {
		case *code.APIError, code.APIError:
			return err
		default:
			logger.Crit("unknown error.", "error", fmt.Sprintf("%+v", err))
			return code.NewAPIError("INTERNAL_SERVER_ERROR")
		}
	}
}
