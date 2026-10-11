package transport_task

// import (
// 	"errors"
// 	"fmt"
// )

// var (
// 	ErrBadRequest = errors.New("400 Bad Request")
// 	// ErrAlreadyExists  = errors.New("already exists")
// 	// ErrInvalidInput   = errors.New("invalid input")
// 	// ErrUnauthorized   = errors.New("unauthorized")
// 	ErrForbidden      = errors.New("403 forbidden")
// 	ErrNorFound       = errors.New("404 page by URL not found")
// 	ErrGatewayTimeout = errors.New("504 Gateway Timeout")
// )

// type TaskResponseError struct {
// 	Err error //?
// 	// StatusCode    int
// 	ClientMessage string
// 	LogMessage    string
// }

// func (re *TaskResponseError) Error() string {
// 	return fmt.Sprintf("ClientMessage: %s, LogMessage: %s", re.ClientMessage, re.LogMessage)
// }

// func GetResponseError(err error) TaskResponseError {
// 	switch {
// 	case errors.Is(err, ErrBadRequest):
// 		return TaskResponseError{
// 			// StatusCode: http.StatusNotFound,
// 			// Code:       ErrNotFound.Error(),
// 			// Message:    "task not found",
// 		}
// 	}
// 	return TaskResponseError{}
// }
