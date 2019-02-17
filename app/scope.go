package app

import (
	"time"
)

type RequestScope interface {
	Logger
	RequestID() string
	Now() time.Time
}
