package errors

import "fmt"

type CustomError struct {
	Message string `json:"error"`
}

func New(title string, e error) *CustomError {
	return &CustomError{
		Message: fmt.Sprintf("[%s]: %s", title, e.Error()),
	}
}

func (e *CustomError) Error() string {
	return e.Message
}
