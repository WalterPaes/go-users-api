package errors

type CustomError struct {
	err error
}

func NewCustomError(e error) *CustomError {
	return &CustomError{
		err: e,
	}
}

func (e *CustomError) Error() string {
	return e.err.Error()
}
