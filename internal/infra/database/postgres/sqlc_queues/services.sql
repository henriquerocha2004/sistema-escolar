-- name: CreateService :exec
INSERT into services (id, description, created_at, updated_at) VALUES ($1,$2,$3,$4);

-- name: UpdateService :exec
UPDATE services SET description = $1, updated_at = $2 WHERE id = $3;

-- name: DeleteService :exec
UPDATE services SET deleted_at = $1 WHERE id = $2;

-- name: FindServiceById :one
SELECT id, description FROM services WHERE id = $1 AND deleted_at IS NULL;