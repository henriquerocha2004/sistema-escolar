package entities

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type Service struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func (s *Service) FillFromDto(dto dto.ServiceRequestDto) {
	s.Description = dto.Description
}
