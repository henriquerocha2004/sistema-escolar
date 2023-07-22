package entities

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type SchoolYear struct {
	Id        uuid.UUID  `json:"id"`
	Year      string     `json:"year"`
	StartedAt *time.Time `json:"started_at"`
	EndAt     *time.Time `json:"end_at"`
	common.Timestamps
}

func (s *SchoolYear) CheckPeriod() error {
	fmt.Println(s)
	if s.EndAt.Before(*s.StartedAt) {
		return errors.New("invalid period provided. EndAt cannot be before that StartedAt")
	}

	return nil
}

func (s *SchoolYear) FillFromDto(dto dto.SchoolYearRequestDto) {
	s.Year = dto.Year
	startedAt, _ := time.Parse("2006-01-02", dto.StartedAt)
	s.StartedAt = &startedAt
	endAt, _ := time.Parse("2006-01-02", dto.EndAt)
	s.EndAt = &endAt
}
