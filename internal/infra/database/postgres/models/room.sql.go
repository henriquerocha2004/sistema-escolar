// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: room.sql

package models

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const bindSchedule = `-- name: BindSchedule :exec
INSERT INTO room_schedule (room_id, schedule_id, school_year_id) VALUES ($1,$2,$3)
`

type BindScheduleParams struct {
	RoomID       uuid.UUID `json:"room_id"`
	ScheduleID   uuid.UUID `json:"schedule_id"`
	SchoolYearID uuid.UUID `json:"school_year_id"`
}

func (q *Queries) BindSchedule(ctx context.Context, arg BindScheduleParams) error {
	_, err := q.db.ExecContext(ctx, bindSchedule, arg.RoomID, arg.ScheduleID, arg.SchoolYearID)
	return err
}

const createRoom = `-- name: CreateRoom :exec
INSERT INTO rooms (id, code, description, capacity, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)
`

type CreateRoomParams struct {
	ID          uuid.UUID    `json:"id"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Capacity    int32        `json:"capacity"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) error {
	_, err := q.db.ExecContext(ctx, createRoom,
		arg.ID,
		arg.Code,
		arg.Description,
		arg.Capacity,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteRoom = `-- name: DeleteRoom :exec
UPDATE rooms SET deleted_at = $1 WHERE id = $2
`

type DeleteRoomParams struct {
	DeletedAt sql.NullTime `json:"deleted_at"`
	ID        uuid.UUID    `json:"id"`
}

func (q *Queries) DeleteRoom(ctx context.Context, arg DeleteRoomParams) error {
	_, err := q.db.ExecContext(ctx, deleteRoom, arg.DeletedAt, arg.ID)
	return err
}

const findByCode = `-- name: FindByCode :one
SELECT id as id, code, description, capacity, created_at FROM rooms WHERE code = $1 AND deleted_at IS NULL
`

type FindByCodeRow struct {
	ID          uuid.UUID    `json:"id"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Capacity    int32        `json:"capacity"`
	CreatedAt   sql.NullTime `json:"created_at"`
}

func (q *Queries) FindByCode(ctx context.Context, code string) (FindByCodeRow, error) {
	row := q.db.QueryRowContext(ctx, findByCode, code)
	var i FindByCodeRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Description,
		&i.Capacity,
		&i.CreatedAt,
	)
	return i, err
}

const findOne = `-- name: FindOne :one
SELECT id as id, code, description, capacity, created_at FROM rooms WHERE id = $1 AND deleted_at IS NULL
`

type FindOneRow struct {
	ID          uuid.UUID    `json:"id"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Capacity    int32        `json:"capacity"`
	CreatedAt   sql.NullTime `json:"created_at"`
}

func (q *Queries) FindOne(ctx context.Context, id uuid.UUID) (FindOneRow, error) {
	row := q.db.QueryRowContext(ctx, findOne, id)
	var i FindOneRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Description,
		&i.Capacity,
		&i.CreatedAt,
	)
	return i, err
}

const unbindSchedule = `-- name: UnbindSchedule :exec
DELETE FROM room_schedule WHERE room_id = $1 AND school_year_id = $2
`

type UnbindScheduleParams struct {
	RoomID       uuid.UUID `json:"room_id"`
	SchoolYearID uuid.UUID `json:"school_year_id"`
}

func (q *Queries) UnbindSchedule(ctx context.Context, arg UnbindScheduleParams) error {
	_, err := q.db.ExecContext(ctx, unbindSchedule, arg.RoomID, arg.SchoolYearID)
	return err
}

const updateRoom = `-- name: UpdateRoom :exec
UPDATE rooms SET code = $1, description = $2, capacity = $3, updated_at = $4 WHERE id = $5
`

type UpdateRoomParams struct {
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Capacity    int32        `json:"capacity"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	ID          uuid.UUID    `json:"id"`
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) error {
	_, err := q.db.ExecContext(ctx, updateRoom,
		arg.Code,
		arg.Description,
		arg.Capacity,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
