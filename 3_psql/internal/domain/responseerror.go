package domain

type ResponseError struct {
	Code ErrorCode
	// Err     error
	Message string
	Log     string
}

func (e *ResponseError) Error() string {
	return e.Log
}

func NewBadRequestError(m string, l string) *ResponseError {
	return &ResponseError{
		Code: ErrBadRequest,
		// Err:     err,
		Message: m,
		Log:     l,
	}
}

func NewNotFoundError(m string, l string) *ResponseError {
	return &ResponseError{
		Code: ErrNotFound,
		// Err:     err,
		Message: m,
		Log:     l,
	}
}

func NewInternalServerError(m string, l string) *ResponseError {
	return &ResponseError{
		Code: ErrInternalServer,
		// Err:     err,
		Message: m,
		Log:     l,
	}
}
