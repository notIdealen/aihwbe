package domain

import "errors"

var (
	ErrNotFound       = errors.New("404 not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("403 forbidden")       // запрос принят, отказано
	ErrGatewayTimeout = errors.New("504 Gateway Timeout") // timeout overflow
)

type ResponseError struct {
	// StatusCode int
	Error   string
	Message string
}
