package dto

import "github.com/go-playground/validator"

type ServiceRequestDto struct {
	Description string `json:"description" validate:"required"`
}

func (s *ServiceRequestDto) Validate() error {
	v := validator.New()
	return v.Struct(s)
}
