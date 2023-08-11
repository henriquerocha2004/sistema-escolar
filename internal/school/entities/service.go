package entities

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type Service struct {
	Id          uuid.UUID
	Description string
}

func (s *Service) FillFromDto(dto dto.ServiceRequestDto) {
	s.Description = dto.Description
}
