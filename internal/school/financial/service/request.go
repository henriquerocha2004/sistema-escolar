package service

import "github.com/go-playground/validator"

type Request struct {
	Description string  `json:"description" validate:"required"`
	Value       float64 `json:"value" validate:"required"`
}

func (s *Request) Validate() error {
	v := validator.New()
	return v.Struct(s)
}
