package validator

import "gopkg.in/go-playground/validator.v9"

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
