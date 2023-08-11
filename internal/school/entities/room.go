package entities

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type Room struct {
	Id          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Capacity    int       `json:"capacity"`
}

func (r *Room) FillFromDto(dto dto.RoomRequestDto) {
	r.Code = dto.Code
	r.Capacity = dto.Capacity
	r.Description = dto.Description
}
