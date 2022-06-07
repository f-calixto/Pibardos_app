package errors

import (
	"errors"
	"fmt"
)

// invalid credentials error
type InvalidCredentials struct {
	Err error
}

func (e *InvalidCredentials) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewInvalidCredentials(msg ...string) *InvalidCredentials {
	defaultMsg := "username or email already in use"
	if len(msg) == 0 {
		return &InvalidCredentials{Err: errors.New(defaultMsg)}
	}
	return &InvalidCredentials{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// Not found error
type NotFound struct {
	Err error
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewNotFound(msg ...string) *NotFound {
	defaultMsg := "user not found"
	if len(msg) == 0 {
		return &NotFound{Err: errors.New(defaultMsg)}
	}
	return &NotFound{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// Invalid update error
type InvalidUpdate struct {
	Err error
}

func (e *InvalidUpdate) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewInvalidUpdate(msg ...string) *InvalidUpdate {
	defaultMsg := "no valid field selected for update"
	if len(msg) == 0 {
		return &InvalidUpdate{Err: errors.New(defaultMsg)}
	}
	return &InvalidUpdate{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// Not found error
type RabbitError struct {
	Err error
}

func (e *RabbitError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewRabbitError(msg ...string) *RabbitError {
	defaultMsg := "rabbitmq unkown error"
	if len(msg) == 0 {
		return &RabbitError{Err: errors.New(defaultMsg)}
	}
	return &RabbitError{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// Not file processing error
type FileError struct {
	Err error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewFileError(msg ...string) *FileError {
	defaultMsg := "error processing file"
	if len(msg) == 0 {
		return &FileError{Err: errors.New(defaultMsg)}
	}
	return &FileError{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// jwt authroization error
type JwtAuthorization struct {
	Err error
}

func (e *JwtAuthorization) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewJwtAuthorization(msg ...string) *JwtAuthorization {
	defaultMsg := "not allowed"
	if len(msg) == 0 {
		return &JwtAuthorization{Err: errors.New(defaultMsg)}
	}
	return &JwtAuthorization{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// jwt bad request
type JwtBadRequest struct {
	Err error
}

func (e *JwtBadRequest) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewJwtBadRequest(msg ...string) *JwtBadRequest {
	defaultMsg := "bad request"
	if len(msg) == 0 {
		return &JwtBadRequest{Err: errors.New(defaultMsg)}
	}
	return &JwtBadRequest{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------

// method not allowed
type MethodNotAllowed struct {
	Err error
}

func (e *MethodNotAllowed) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewMethodNotAllowed(msg ...string) *MethodNotAllowed {
	defaultMsg := "method not allowed"
	if len(msg) == 0 {
		return &MethodNotAllowed{Err: errors.New(defaultMsg)}
	}
	return &MethodNotAllowed{Err: errors.New(msg[0])}
}

// ---
