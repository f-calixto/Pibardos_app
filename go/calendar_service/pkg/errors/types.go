package errors

import (
	"errors"
	"fmt"
)

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

// -------------------------------------------------------------

// already joined event
type AlreadyJoined struct {
	Err error
}

func (e *AlreadyJoined) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewAlreadyJoined(msg ...string) *AlreadyJoined {
	defaultMsg := "already joined event"
	if len(msg) == 0 {
		return &AlreadyJoined{Err: errors.New(defaultMsg)}
	}
	return &AlreadyJoined{Err: errors.New(msg[0])}
}

// -------------------------------------------------------------

// already joined event
type InvalidDate struct {
	Err error
}

func (e *InvalidDate) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewInvalidDate(msg ...string) *InvalidDate {
	defaultMsg := "invalid dates"
	if len(msg) == 0 {
		return &InvalidDate{Err: errors.New(defaultMsg)}
	}
	return &InvalidDate{Err: errors.New(msg[0])}
}

// -------------------------------------------------------------

// invalid new event date
type OverlapingDate struct {
	Err error
}

func (e *OverlapingDate) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewOverlapingDate(msg ...string) *OverlapingDate {
	defaultMsg := "invalid dates"
	if len(msg) == 0 {
		return &OverlapingDate{Err: errors.New(defaultMsg)}
	}
	return &OverlapingDate{Err: errors.New(msg[0])}
}

// -------------------------------------------------------------

// not authorized to do x
type Unauthorized struct {
	Err error
}

func (e *Unauthorized) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewUnauthorized(msg ...string) *Unauthorized {
	defaultMsg := "not authorized"
	if len(msg) == 0 {
		return &Unauthorized{Err: errors.New(defaultMsg)}
	}
	return &Unauthorized{Err: errors.New(msg[0])}
}

// -------------------------------------------------------------

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
