-- name: CreateService :exec
INSERT into services (id, description, price, created_at, updated_at) VALUES ($1,$2,$3,$4,$5);

-- name: UpdateService :exec
UPDATE services SET description = $1, price = $2, updated_at = $3 WHERE id = $4;

-- name: DeleteService :exec
UPDATE services SET deleted_at = $1 WHERE id = $2;

-- name: FindServiceById :one
SELECT id, description, price FROM services WHERE id = $1 AND deleted_at IS NULL;