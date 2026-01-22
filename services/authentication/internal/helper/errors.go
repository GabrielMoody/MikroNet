package helper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = fmt.Errorf("data not found")
	ErrDuplicateEntry    = fmt.Errorf("duplicate entry on email")
	ErrDatabase          = fmt.Errorf("database error")
	ErrBadRequest        = fmt.Errorf("bad request")
	ErrPasswordIncorrect = fmt.Errorf("password incorrect")
	ErrBlockedAccount    = fmt.Errorf("akun anda telah diblokir")
	ErrExpired           = fmt.Errorf("link reset password telah expired/invalid. silahkan melakukan reset password kembali")
	ErrNotVerified       = fmt.Errorf("akun anda belum diverifikasi")
)

type ErrorStruct struct {
	Err              error
	Code             int
	ValidationErrors map[string]string
}

var errorMessages = map[string]string{
	"required": "is required",
	"email":    "must be a valid email address",
	"min":      "must be greater than %s characters",
	"eqfield":  "must be the same with %s",
}

func translateError(err validator.FieldError) string {
	msg, ok := errorMessages[err.Tag()]
	if !ok {
		return err.Tag()
	}

	if err.Param() != "" {
		return fmt.Sprintf(msg, err.Param())
	}
	return msg
}

func ValidationError(err error) map[string]string {
	errField := make(map[string]string)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errField[err.Field()] = fmt.Sprintf("%s %s", err.Field(), translateError(err))
		}
	}
	return errField
}

func CheckError(err error) *ErrorStruct {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ErrNotFound):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusNotFound,
		}
	case errors.Is(err, ErrDuplicateEntry):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusConflict,
		}
	case errors.Is(err, ErrDatabase):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	case errors.Is(err, ErrBadRequest):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	case errors.Is(err, ErrPasswordIncorrect):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	case errors.Is(err, ErrBlockedAccount):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusForbidden,
		}
	case errors.Is(err, ErrNotVerified):
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusForbidden,
		}
	default:
		return &ErrorStruct{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}
}

var Validate = validator.New()
