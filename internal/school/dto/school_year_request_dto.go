package dto

import (
	"github.com/go-playground/validator"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
)

type SchoolYearRequestDto struct {
	Id        string `json:"id,omitempty"`
	Year      string `json:"year" validate:"required"`
	StartedAt string `json:"start_at" validate:"required,date::format:yyyy-mm-dd"`
	EndAt     string `json:"end_at" validate:"required,date::format:yyyy-mm-dd"`
}

func (s *SchoolYearRequestDto) Validate() error {
	v := validator.New()
	_ = v.RegisterValidation("date::format:yyyy-mm-dd", requestvalidator.ValidateDateUSA)
	return v.Struct(s)
}
