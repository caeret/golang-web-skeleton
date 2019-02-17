package app

import (
	"fmt"
	"net/http"
	"time"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/fault"
)

func RoutePrepare(logger Logger) routing.Handler {
	return func(c *routing.Context) error {
		now := time.Now()
		c.Response = &access.LogResponseWriter{c.Response, http.StatusOK, 0}
		rs := newRequestScope(now, logger, c.Request)
		c.Set("Context", rs)
		fault.Recovery(func(format string, a ...interface{}) {
			rs.Crit(fmt.Sprintf(format, a...))
		}, convertError)(c)
		logAccess(c, func(format string, a ...interface{}) {
			rs.Info(fmt.Sprintf(format, a...))
		}, rs.Now())
		return nil
	}
}

func GetRequestScope(c *routing.Context) RequestScope {
	return c.Get("Context").(RequestScope)
}

func logAccess(c *routing.Context, logFunc access.LogFunc, start time.Time) {
	rw := c.Response.(*access.LogResponseWriter)
	elapsed := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	requestLine := fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto)
	logFunc(`%.3fms %s %d %d`, elapsed, requestLine, rw.Status, rw.BytesWritten)
}

func newRequestScope(now time.Time, logger Logger, request *http.Request) RequestScope {
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		logger = logger.New("request_id", requestID)
	}
	return &requestScope{
		Logger:    logger,
		now:       now,
		requestID: requestID,
	}
}

func convertError(c *routing.Context, err error) error {
	return err
}
