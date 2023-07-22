// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: school_year.sql

package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createYearSchool = `-- name: CreateYearSchool :exec
INSERT INTO school_year (id,year,start_at,end_at,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6)
`

type CreateYearSchoolParams struct {
	ID        uuid.UUID    `json:"id"`
	Year      string       `json:"year"`
	StartAt   time.Time    `json:"start_at"`
	EndAt     time.Time    `json:"end_at"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateYearSchool(ctx context.Context, arg CreateYearSchoolParams) error {
	_, err := q.db.ExecContext(ctx, createYearSchool,
		arg.ID,
		arg.Year,
		arg.StartAt,
		arg.EndAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteYearSchool = `-- name: DeleteYearSchool :exec
UPDATE school_year SET deleted_at = $1 WHERE id = $2
`

type DeleteYearSchoolParams struct {
	DeletedAt sql.NullTime `json:"deleted_at"`
	ID        uuid.UUID    `json:"id"`
}

func (q *Queries) DeleteYearSchool(ctx context.Context, arg DeleteYearSchoolParams) error {
	_, err := q.db.ExecContext(ctx, deleteYearSchool, arg.DeletedAt, arg.ID)
	return err
}

const findByYear = `-- name: FindByYear :one
SELECT id, year, start_at, end_at FROM school_year WHERE year = $1 LIMIT 1
`

type FindByYearRow struct {
	ID      uuid.UUID `json:"id"`
	Year    string    `json:"year"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func (q *Queries) FindByYear(ctx context.Context, year string) (FindByYearRow, error) {
	row := q.db.QueryRowContext(ctx, findByYear, year)
	var i FindByYearRow
	err := row.Scan(
		&i.ID,
		&i.Year,
		&i.StartAt,
		&i.EndAt,
	)
	return i, err
}

const findOneSchoolYear = `-- name: FindOneSchoolYear :one
SELECT id, year, start_at, end_at FROM school_year WHERE id = $1 AND deleted_at IS NULL
`

type FindOneSchoolYearRow struct {
	ID      uuid.UUID `json:"id"`
	Year    string    `json:"year"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func (q *Queries) FindOneSchoolYear(ctx context.Context, id uuid.UUID) (FindOneSchoolYearRow, error) {
	row := q.db.QueryRowContext(ctx, findOneSchoolYear, id)
	var i FindOneSchoolYearRow
	err := row.Scan(
		&i.ID,
		&i.Year,
		&i.StartAt,
		&i.EndAt,
	)
	return i, err
}

const updateSchoolYear = `-- name: UpdateSchoolYear :exec
UPDATE school_year SET year = $1, start_at = $2, end_at = $3, updated_at = $4 WHERE id = $5
`

type UpdateSchoolYearParams struct {
	Year      string       `json:"year"`
	StartAt   time.Time    `json:"start_at"`
	EndAt     time.Time    `json:"end_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ID        uuid.UUID    `json:"id"`
}

func (q *Queries) UpdateSchoolYear(ctx context.Context, arg UpdateSchoolYearParams) error {
	_, err := q.db.ExecContext(ctx, updateSchoolYear,
		arg.Year,
		arg.StartAt,
		arg.EndAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
