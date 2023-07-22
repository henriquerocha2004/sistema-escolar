-- name: CreateSchedule :exec
INSERT INTO class_schedule (id, description, schedule, school_year_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6);

-- name: DeleteSchedule :exec
UPDATE class_schedule SET deleted_at = $1 WHERE id = $2;

-- name: UpdateSchedule :exec
UPDATE class_schedule SET description = $1, schedule = $2, school_year_id = $3, updated_at = $4 WHERE id = $5;

-- name: FindOneSchedule :one
SELECT class_schedule.id as schedule_id, description, schedule, school_year.id FROM class_schedule
     JOIN school_year ON school_year.id = class_schedule.school_year_id
     WHERE class_schedule.id = $1 AND class_schedule.deleted_at IS NULL;
     
-- name: FindBySchoolYearId :many
SELECT class_schedule.id as schedule_id, description, schedule, school_year.year FROM class_schedule
     JOIN school_year ON school_year.id = class_schedule.school_year_id
     WHERE class_schedule.school_year_id = $1 AND class_schedule.deleted_at IS NULL;