package utils

import (
	"fmt"
	"net/http"
)

type Error interface {
	GetCode() int
	GetDetail() string
}

type httpError struct {
	Detail string `json:"detail"`
	Code   int    `json:"code"`
}

func (e httpError) GetCode() int {
	return e.Code
}

func (e httpError) GetDetail() string {
	return e.Detail
}

func NewInternal(detail string) Error {
	return httpError{
		Detail: detail,
		Code:   http.StatusInternalServerError,
	}
}

func NewInternalf(template string, args ...interface{}) Error {
	return httpError{
		Detail: fmt.Sprintf(template, args...),
		Code:   http.StatusInternalServerError,
	}
}

func NewNotFound(detail string) Error {
	return httpError{
		Detail: detail,
		Code:   http.StatusNotFound,
	}
}

func NewBadRequest(detail string) Error {
	return httpError{
		Detail: detail,
		Code:   http.StatusBadRequest,
	}
}
