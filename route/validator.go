package route

import "github.com/go-playground/validator/v10"

// Validator strucg
type Validator struct {
	validator *validator.Validate
}

// NewValidator new validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate struct
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
