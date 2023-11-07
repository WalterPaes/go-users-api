package errors

import "fmt"

type CustomError struct {
	title string
	err   error
}

func New(title string, e error) *CustomError {
	return &CustomError{
		title: title,
		err:   e,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.title, e.err.Error())
}
