-- name: CreateRoom :exec
INSERT INTO rooms (id, code, description, capacity, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6);

-- name: DeleteRoom :exec
UPDATE rooms SET deleted_at = $1 WHERE id = $2;

-- name: UpdateRoom :exec
UPDATE rooms SET code = $1, description = $2, capacity = $3, updated_at = $4 WHERE id = $5;

-- name: FindOne :one
SELECT id as id, code, description, capacity, created_at FROM rooms WHERE id = $1 AND deleted_at IS NULL;

-- name: FindByCode :one
SELECT id as id, code, description, capacity, created_at FROM rooms WHERE code = $1 AND deleted_at IS NULL;

-- name: UnbindSchedule :exec
DELETE FROM room_schedule WHERE room_id = $1 AND school_year_id = $2;

-- name: BindSchedule :exec
INSERT INTO room_schedule (room_id, schedule_id, school_year_id) VALUES ($1,$2,$3);