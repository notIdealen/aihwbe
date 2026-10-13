package service_error

import "errors"

type ServiceError struct {
	C ErrorCode
	M string
	L string
}

func (se *ServiceError) Error() string {
	return se.L
}

// func NewStorageError(m, l string) error {
// 	return &ServiceError{
// 		C: ErrInternalServer,
// 		M: m,
// 		L: l,
// 	}
// }

func NewValidationError(m, l string) error {
	return &ServiceError{
		C: ErrValidation,
		M: m,
		L: l,
	}
}

func IsServiceError(err error) (*ServiceError, bool) {
	if se, ok := errors.AsType[*ServiceError](err); ok {
		return se, true
	}
	return nil, false
}
