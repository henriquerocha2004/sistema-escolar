package dto

import (
	"github.com/go-playground/validator"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
)

type ScheduleRequestDto struct {
	Description string `json:"description" validate:"omitempty"`
	InitialTime string `json:"initial_time" validate:"required,time"`
	FinalTime   string `json:"final_time" validate:"required,time"`
	SchoolYear  string `json:"school_year" validate:"required"`
}

func (s *ScheduleRequestDto) Validate() error {
	v := validator.New()
	_ = v.RegisterValidation("time", requestvalidator.ValidateHourFormat)
	return v.Struct(s)
}
