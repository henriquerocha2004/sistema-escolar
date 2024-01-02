package entities

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type ScheduleClass struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Schedule    string    `json:"schedule"`
	SchoolYear  uuid.UUID `json:"school_year"`
}

func (s *ScheduleClass) FillFromDto(dto dto.ScheduleRequestDto) error {
	s.Description = dto.Description
	err := s.setSchedule(dto.InitialTime, dto.FinalTime)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ScheduleClass) setSchedule(initialTime string, finalTime string) error {
	start := fmt.Sprintf("%s:00", initialTime)
	end := fmt.Sprintf("%s:00", finalTime)

	t1, err := time.Parse("15:04:05", start)
	if err != nil {
		return err
	}
	t2, err := time.Parse("15:04:05", end)
	if err != nil {
		return err
	}

	diffMinutes := t1.Sub(t2).Minutes()
	if diffMinutes > 0 {
		return errors.New("initial time can not be greather than final time")
	}

	s.Schedule = fmt.Sprintf("%s-%s", initialTime, finalTime)

	return nil
}
