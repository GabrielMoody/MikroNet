package helper

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = fmt.Errorf("data not found")
	ErrDuplicateEntry    = fmt.Errorf("duplicate entry on email")
	ErrDatabase          = fmt.Errorf("database error")
	ErrBadRequest        = fmt.Errorf("bad request")
	ErrPasswordIncorrect = fmt.Errorf("password incorrect")
	ErrBlocked           = fmt.Errorf("user blocked")
)

type ErrorStruct struct {
	Err  error
	Code int
}

func CheckError(err error) *ErrorStruct {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ErrNotFound):
		return &ErrorStruct{
			Err:  err,
			Code: 404,
		}
	case errors.Is(err, ErrDuplicateEntry):
		return &ErrorStruct{
			Err:  err,
			Code: 409,
		}
	case errors.Is(err, ErrDatabase):
		return &ErrorStruct{
			Err:  err,
			Code: 500,
		}
	case errors.Is(err, ErrBadRequest):
		return &ErrorStruct{
			Err:  err,
			Code: 400,
		}
	case errors.Is(err, ErrPasswordIncorrect):
		return &ErrorStruct{
			Err:  err,
			Code: 401,
		}
	case errors.Is(err, ErrBlocked):
		return &ErrorStruct{
			Err:  err,
			Code: 403,
		}
	default:
		return &ErrorStruct{
			Err:  err,
			Code: 500,
		}
	}
}

var Validate = validator.New()
