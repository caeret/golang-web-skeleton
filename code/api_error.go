package code

import "net/http"

func NewAPIError(code string) *APIError {
	err := &APIError{
		Status:    http.StatusBadRequest,
		ErrorCode: code,
		Message:   code,
	}
	return err
}

type APIError struct {
	Status           int         `json:"-"`
	ErrorCode        string      `json:"code"`
	Message          string      `json:"message"`
	DeveloperMessage string      `json:"developer_message,omitempty"`
	Details          interface{} `json:"details,omitempty"`
}

func (e *APIError) WithDetails(d interface{}) *APIError {
	e.Details = d
	return e
}

func (e APIError) Error() string {
	return e.Message
}

func (e APIError) StatusCode() int {
	return e.Status
}
