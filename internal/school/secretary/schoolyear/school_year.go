package schoolyear

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type SchoolYear struct {
	id        uuid.UUID
	year      string
	startedAt *time.Time
	endAt     *time.Time
}

func New(year string, startAt string, endAt string) (*SchoolYear, error) {

	sy := &SchoolYear{
		id: uuid.New(),
	}

	err := sy.ChangeSchoolYear(year)
	if err != nil {
		return nil, err
	}

	err = sy.ChangeStartAt(startAt)
	if err != nil {
		return nil, err
	}

	err = sy.ChangeEndAt(endAt)
	if err != nil {
		return nil, err
	}

	err = sy.CheckPeriod()
	if err != nil {
		return nil, err
	}

	return sy, nil
}

func (sy *SchoolYear) ChangeSchoolYear(year string) error {
	if year == "" {
		return errors.New("year cannot be empty")
	}

	sy.year = year

	return nil
}

func (sy *SchoolYear) ChangeStartAt(startAt string) error {
	sa, err := time.Parse("2006-01-02", startAt)

	if err != nil {
		return err
	}

	sy.startedAt = &sa

	return nil
}

func (sy *SchoolYear) ChangeEndAt(endAt string) error {
	et, err := time.Parse("2006-01-02", endAt)

	if err != nil {
		return err
	}

	sy.endAt = &et

	return nil
}

func (sy *SchoolYear) CheckPeriod() error {
	if sy.endAt.Before(*sy.startedAt) {
		return errors.New("invalid period provided. EndAt cannot be before that StartedAt")
	}

	return nil
}

func (sy *SchoolYear) SetId(id string) error {
	schoolYearId, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	sy.id = schoolYearId

	return nil
}

func (sy *SchoolYear) Year() string {
	return sy.year
}

func (sy *SchoolYear) StartAt() *time.Time {
	return sy.startedAt
}

func (sy *SchoolYear) EndAt() *time.Time {
	return sy.endAt
}

func (sy *SchoolYear) Id() uuid.UUID {
	return sy.id
}

func (sy *SchoolYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id      string `json:"id"`
		Year    string `json:"year"`
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
	}{
		Id:      sy.id.String(),
		Year:    sy.Year(),
		StartAt: sy.StartAt().Format("2006-01-02"),
		EndAt:   sy.EndAt().Format("2006-01-02"),
	})
}
