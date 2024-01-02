-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
-- name: CreateAddress :exec
INSERT INTO addresses
(id, street, city, district, state, zip_code, owner_id, created_at, updated_at)
VALUES
($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: DeleteAddressByOwner :exec
UPDATE addresses SET deleted_at = $1 WHERE owner_id = $2;