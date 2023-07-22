package entities

import (
	"errors"
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"time"
)

type ClassRoom struct {
	Id              uuid.UUID     `json:"id"`
	VacancyQuantity int           `json:"vacancy_quantity"`
	Shift           string        `json:"shift"`
	OpenDate        *time.Time    `json:"open_date"`
	OccupiedVacancy int           `json:"occupied_vacancy"`
	Status          string        `json:"status"`
	Level           string        `json:"level"`
	Identification  string        `json:"identification"`
	SchoolYearId    uuid.UUID     `json:"school_year_id"`
	RoomId          uuid.NullUUID `json:"room_id"`
	ScheduleId      uuid.UUID     `json:"schedule_id"`
	Localization    string        `json:"localization"`
	Type            string        `json:"type"`
	common.Timestamps
}

func (c *ClassRoom) SetOccupiedVacancies(quantity int) error {

	err := c.CheckStatus()
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return nil
	}

	remainingVacancies := c.VacancyQuantity - c.OccupiedVacancy

	if quantity > remainingVacancies {
		return errors.New("quantidade de vagas disponíveis é inferior a quantidade de vagas solicitadas")
	}

	if quantity > c.VacancyQuantity {
		return errors.New("quantidade de vagas para essa classe é inferior a quantidade solicitada")
	}

	c.OccupiedVacancy += quantity
	return nil
}

func (c *ClassRoom) GetOccupiedVacancies() int {
	return c.OccupiedVacancy
}

func (c *ClassRoom) CheckStatus() error {
	if c.Status == "closed" {
		return errors.New("class Room is closed")
	}

	return nil
}

func (c *ClassRoom) FillFromDto(dto dto.ClassRoomRequestDto) {
	scheduleId, _ := uuid.Parse(dto.ScheduleId)
	roomId, _ := uuid.Parse(dto.RoomId)
	schoolYearId, _ := uuid.Parse(dto.SchoolYearId)
	openDate := time.Now()

	c.Type = dto.Type
	c.Localization = dto.Localization
	c.Shift = dto.Shift
	c.ScheduleId = scheduleId
	c.OpenDate = &openDate
	c.RoomId = uuid.NullUUID{
		UUID:  roomId,
		Valid: true,
	}
	c.SchoolYearId = schoolYearId
	c.Level = dto.Level
	c.VacancyQuantity = dto.VacancyQuantity
	c.Identification = dto.Identification
	c.Status = dto.Status
}
