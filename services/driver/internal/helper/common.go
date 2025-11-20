package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound = fmt.Errorf("data not found")
	ErrDatabase = fmt.Errorf("database error")
)

type ErrorStruct struct {
	Err  error
	Code int
}

var Validate = validator.New()
