package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-routing/fault"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
)

type RequestScope interface {
	Logger
	Container
	RequestID() string
	Now() time.Time
}

type requestScope struct {
	Logger
	Container
	now       time.Time
	requestID string
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func RoutePrepare(logger Logger, container Container) routing.Handler {
	return func(rc *routing.Context) error {
		now := time.Now()
		rc.Response = &access.LogResponseWriter{rc.Response, http.StatusOK, 0}
		ac := newRequestScope(now, logger, rc.Request, container)
		rc.Set("Context", ac)
		fault.Recovery(func(format string, a ...interface{}) {
			ac.Crit(fmt.Sprintf(format, a...))
		}, convertError)(rc)
		logAccess(rc, func(format string, a ...interface{}) {
			ac.Info(fmt.Sprintf(format, a...))
		}, ac.Now())
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
	logFunc(`[%.3fms] %s %d %d`, elapsed, requestLine, rw.Status, rw.BytesWritten)
}

func newRequestScope(now time.Time, logger Logger, request *http.Request, container Container) RequestScope {
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		logger = logger.New("request_id", requestID)
	}
	return &requestScope{
		Logger:    logger,
		Container: container,
		now:       now,
		requestID: requestID,
	}
}

func convertError(c *routing.Context, err error) error {
	return err
}
