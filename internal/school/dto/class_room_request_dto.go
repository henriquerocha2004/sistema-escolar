package dto

import (
	"github.com/go-playground/validator"
	requestValidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
)

type ClassRoomRequestDto struct {
	VacancyQuantity int    `json:"vacancy_quantity" validate:"required,number"`
	Shift           string `json:"shift" validate:"required,shift"`
	Status          string `json:"status" validate:"required,status"`
	Level           string `json:"level" validate:"required"`
	Identification  string `json:"identification" validate:"required"`
	SchoolYearId    string `json:"school_year_id" validate:"required,uuid"`
	RoomId          string `json:"room_id" validate:"omitempty,uuid"`
	ScheduleId      string `json:"schedule_id" validate:"required,uuid"`
	Localization    string `json:"localization" validate:"required"`
	Type            string `json:"type" validate:"required,type"`
}

func (c *ClassRoomRequestDto) Validate() error {
	v := validator.New()
	_ = v.RegisterValidation("shift", requestValidator.ValidateShift)
	_ = v.RegisterValidation("type", requestValidator.ValidateType)
	_ = v.RegisterValidation("status", requestValidator.ValidateClassStatus)

	return v.Struct(c)
}
