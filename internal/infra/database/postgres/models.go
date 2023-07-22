// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ClassRoom struct {
	ID                uuid.UUID      `json:"id"`
	Active            bool           `json:"active"`
	Status            string         `json:"status"`
	Identification    string         `json:"identification"`
	Type              string         `json:"type"`
	Vacancies         int32          `json:"vacancies"`
	VacanciesOccupied int32          `json:"vacancies_occupied"`
	Shift             string         `json:"shift"`
	Level             string         `json:"level"`
	Localization      sql.NullString `json:"localization"`
	OpenDate          time.Time      `json:"open_date"`
	SchoolYearID      uuid.UUID      `json:"school_year_id"`
	RoomID            uuid.NullUUID  `json:"room_id"`
	ScheduleID        uuid.UUID      `json:"schedule_id"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
}

type ClassSchedule struct {
	ID           uuid.UUID    `json:"id"`
	Description  string       `json:"description"`
	Schedule     string       `json:"schedule"`
	SchoolYearID uuid.UUID    `json:"school_year_id"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type Room struct {
	ID          uuid.UUID    `json:"id"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Capacity    int32        `json:"capacity"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type RoomSchedule struct {
	ID           int64     `json:"id"`
	RoomID       uuid.UUID `json:"room_id"`
	ScheduleID   uuid.UUID `json:"schedule_id"`
	SchoolYearID uuid.UUID `json:"school_year_id"`
}

type SchoolYear struct {
	ID        uuid.UUID    `json:"id"`
	Year      string       `json:"year"`
	StartAt   time.Time    `json:"start_at"`
	EndAt     time.Time    `json:"end_at"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}