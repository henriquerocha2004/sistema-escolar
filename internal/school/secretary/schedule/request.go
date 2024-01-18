package schedule

import (
	"github.com/go-playground/validator"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
)

type Request struct {
	Description string `json:"description" validate:"omitempty"`
	InitialTime string `json:"initial_time" validate:"required,time"`
	FinalTime   string `json:"final_time" validate:"required,time"`
	SchoolYear  string `json:"school_year" validate:"required"`
}

func (s *Request) Validate() error {
	v := validator.New()
	_ = v.RegisterValidation("time", requestvalidator.ValidateHourFormat)
	return v.Struct(s)
}
