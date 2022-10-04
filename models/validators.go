package models

import "github.com/go-playground/validator/v10"

var (
	V = validator.New()
)

type PersonValidator struct {
	Validator *validator.Validate
}

func (p *PersonValidator) Validate(i interface{}) error {
	return p.Validator.Struct(i)
}
