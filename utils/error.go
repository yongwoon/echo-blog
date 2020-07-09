package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Error struct
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewError new error
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["message"] = v.Message
	default:
		e.Errors["message"] = v.Error()
	}
	return e
}

// NewValidatorError validation error
func NewValidatorError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

// AccessForbidden 403 forbidden error
func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = "access forbidden"
	return e
}

// NotFound 404 not found error
func NotFound() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = "resource not found"
	return e
}
