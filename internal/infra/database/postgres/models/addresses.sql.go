// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: addresses.sql

package models

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAddress = `-- name: CreateAddress :exec
INSERT INTO addresses
(id, street, city, district, state, zip_code, owner_id, created_at, updated_at)
VALUES
($1,$2,$3,$4,$5,$6,$7,$8,$9)
`

type CreateAddressParams struct {
	ID        uuid.UUID    `json:"id"`
	Street    string       `json:"street"`
	City      string       `json:"city"`
	District  string       `json:"district"`
	State     string       `json:"state"`
	ZipCode   string       `json:"zip_code"`
	OwnerID   uuid.UUID    `json:"owner_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

// Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) error {
	_, err := q.db.ExecContext(ctx, createAddress,
		arg.ID,
		arg.Street,
		arg.City,
		arg.District,
		arg.State,
		arg.ZipCode,
		arg.OwnerID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteAddressByOwner = `-- name: DeleteAddressByOwner :exec
UPDATE addresses SET deleted_at = $1 WHERE owner_id = $2
`

type DeleteAddressByOwnerParams struct {
	DeletedAt sql.NullTime `json:"deleted_at"`
	OwnerID   uuid.UUID    `json:"owner_id"`
}

func (q *Queries) DeleteAddressByOwner(ctx context.Context, arg DeleteAddressByOwnerParams) error {
	_, err := q.db.ExecContext(ctx, deleteAddressByOwner, arg.DeletedAt, arg.OwnerID)
	return err
}