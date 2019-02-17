package scope

import (
	"net/http"
	"time"

	"github.com/caeret/golang-web-skeleton/app"
	routing "github.com/go-ozzo/ozzo-routing"
)

type requestScope struct {
	app.Logger
	now       time.Time
	requestID string
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func GetRequestScope(c *routing.Context) app.RequestScope {
	return c.Get("Context").(app.RequestScope)
}

func NewRequestScope(now time.Time, logger app.Logger, request *http.Request) app.RequestScope {
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
