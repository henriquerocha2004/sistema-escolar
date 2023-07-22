package dto

import "github.com/go-playground/validator"

type RoomScheduleDto struct {
	SchoolYear  string   `json:"school_year" validate:"required,uuid"`
	RoomId      string   `json:"room_id" validate:"required,uuid"`
	ScheduleIds []string `json:"schedule_ids" validate:"required,dive,uuid"`
}

func (r *RoomScheduleDto) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
