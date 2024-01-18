package schedule

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type ScheduleClass struct {
	id          uuid.UUID
	description string
	startAt     string
	endAt       string
	schoolYear  uuid.UUID
}

func New(description string, initialTime string, finalTime string, schoolYearId string) (*ScheduleClass, error) {
	s := &ScheduleClass{
		id: uuid.New(),
	}

	err := s.ChangeDescription(description)
	if err != nil {
		return nil, err
	}

	err = s.ChangePeriod(initialTime, finalTime)
	if err != nil {
		return nil, err
	}

	err = s.ChangeSchoolYearId(schoolYearId)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func Load(id string, description string, initialTime string, finalTime string, schoolYearId string) (*ScheduleClass, error) {
	schedule, err := New(description, initialTime, finalTime, schoolYearId)
	if err != nil {
		return nil, err
	}

	_ = schedule.ChangeId(id)

	return schedule, nil
}

func (s *ScheduleClass) Description() string {
	return s.description
}

func (s *ScheduleClass) StartAt() string {
	return s.startAt
}

func (s *ScheduleClass) EndAt() string {
	return s.endAt
}

func (s *ScheduleClass) SchoolYearId() uuid.UUID {
	return s.schoolYear
}

func (s *ScheduleClass) Id() uuid.UUID {
	return s.id
}

func (s *ScheduleClass) ChangeId(id string) error {
	if id == "" {
		return errors.New("schedule id cannot be empty")
	}

	schedule, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed do change schedule id")
	}

	s.id = schedule

	return nil
}

func (s *ScheduleClass) ChangeSchoolYearId(schoolYearId string) error {
	if schoolYearId == "" {
		return errors.New("school year id cannot be empty")
	}

	schoolYear, err := uuid.Parse(schoolYearId)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change school year id")
	}

	s.schoolYear = schoolYear

	return nil
}

func (s *ScheduleClass) ChangeDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}

	s.description = description

	return nil
}

func (s *ScheduleClass) ChangePeriod(initialTime string, finalTime string) error {

	err := s.validateTime(initialTime)
	if err != nil {
		return err
	}

	err = s.validateTime(finalTime)
	if err != nil {
		return err
	}

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
		return errors.New("initial time can not be greater than final time")
	}

	s.startAt = start
	s.endAt = end

	return nil
}

func (s *ScheduleClass) validateTime(hour string) error {
	regex, _ := regexp.Compile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$")
	if !regex.MatchString(hour) {
		return errors.New("invalid time schedule provided")
	}

	return nil
}

func (s *ScheduleClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id          string `json:"id"`
		Description string `json:"description"`
		StartAt     string `json:"start_at"`
		EndAt       string `json:"end_at"`
		SchoolYear  string `json:"school_year_id"`
	}{
		Id:          s.Id().String(),
		Description: s.Description(),
		StartAt:     s.StartAt(),
		EndAt:       s.EndAt(),
		SchoolYear:  s.SchoolYearId().String(),
	})
}
