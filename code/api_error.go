package code

func NewAPIError(code string) *APIError {
	err := &APIError{
		Code:    code,
		Message: code,
	}
	return err
}

type APIError struct {
	Code             string      `json:"code"`
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
