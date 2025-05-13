package models

import (
	"errors"
	"fmt"
	"net/http"
)

type RequestError struct {
	Code     int
	ErrorMsg error
}

type ErrorResponse struct {
	Title   string
	Message string
	Code    int
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status: %d\nmessage: %v", r.Code, r.ErrorMsg)
}

func UnauthorizedError() error {
	return &RequestError{
		Code:     http.StatusUnauthorized,
		ErrorMsg: errors.New("UNAUTHORIZED"),
	}
}

func ForbiddenError() error {
	return &RequestError{
		Code:     http.StatusForbidden,
		ErrorMsg: errors.New("FORBIDDEN"),
	}
}

func NotFoundError() error {
	return &RequestError{
		Code:     http.StatusNotFound,
		ErrorMsg: errors.New("NOT_FOUND"),
	}
}

func InternalServerError() error {
	return &RequestError{
		Code:     http.StatusInternalServerError,
		ErrorMsg: errors.New("INTERNAL_SERVER_ERROR"),
	}
}
