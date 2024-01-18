-- name: CreateSchedule :exec
INSERT INTO class_schedule (id, description, start_at, end_at, school_year_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7);

-- name: DeleteSchedule :exec
UPDATE class_schedule SET deleted_at = $1 WHERE id = $2;

-- name: UpdateSchedule :exec
UPDATE class_schedule SET description = $1, start_at = $2, end_at = $3, school_year_id = $4, updated_at = $5 WHERE id = $6;

-- name: FindOneSchedule :one
SELECT class_schedule.id as schedule_id, description, class_schedule.start_at, class_schedule.end_at, school_year.id FROM class_schedule
     JOIN school_year ON school_year.id = class_schedule.school_year_id
     WHERE class_schedule.id = $1 AND class_schedule.deleted_at IS NULL;
     
-- name: FindBySchoolYearId :many
SELECT class_schedule.id as schedule_id, description, class_schedule.start_at, class_schedule.end_at, school_year.year FROM class_schedule
     JOIN school_year ON school_year.id = class_schedule.school_year_id
     WHERE class_schedule.school_year_id = $1 AND class_schedule.deleted_at IS NULL;