package rerr

import "net/http"

// RestError interface
type RestError interface {
	Message() string
	Status() int
	Title() string
}

type restError struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
	ErrorTitle   string `json:"title"`
}

func (e restError) Message() string {
	return e.ErrorMessage
}

func (e restError) Status() int {
	return e.ErrorStatus
}

func (e restError) Title() string {
	return e.ErrorTitle
}

// NewBadRequestError return a RestError interface
func NewBadRequestError(message string) RestError {
	return restError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusBadRequest,
		ErrorTitle:   "bad_request",
	}
}

// NewNotFoundError return a RestError interface
func NewNotFoundError(message string) RestError {
	return restError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusNotFound,
		ErrorTitle:   "not_found",
	}
}
