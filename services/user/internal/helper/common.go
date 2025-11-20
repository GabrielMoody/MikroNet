package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = fmt.Errorf("data not found")
	ErrDuplicateEntry    = fmt.Errorf("duplicate entry on email")
	ErrInvalidInput      = fmt.Errorf("invalid input")
	ErrUnauthorized      = fmt.Errorf("unauthorized")
	ErrDatabase          = fmt.Errorf("database error")
	ErrInternal          = fmt.Errorf("internal server error")
	ErrBadRequest        = fmt.Errorf("bad request")
	ErrPasswordIncorrect = fmt.Errorf("password incorrect")
)

type ErrorStruct struct {
	Err  error
	Code int
}

var Validate = validator.New()
