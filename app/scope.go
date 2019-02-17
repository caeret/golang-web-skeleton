package app

import (
	"time"
)

type RequestScope interface {
	Logger
	RequestID() string
	Now() time.Time
}

type requestScope struct {
	Logger
	now       time.Time
	requestID string
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}
