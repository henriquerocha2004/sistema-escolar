package dto

import (
	"github.com/go-playground/validator"
)

type RoomRequestDto struct {
	Code        string   `json:"code" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Capacity    int      `json:"capacity" validate:"required"`
	Schedules   []string `json:"schedules,omitempty" validate:"omitempty,dive,uuid"`
}

func (r *RoomRequestDto) Validate() error {
	validator := validator.New()
	return validator.Struct(r)
}
