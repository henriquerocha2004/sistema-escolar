-- name: CreateYearSchool :exec
INSERT INTO school_year (id,year,start_at,end_at,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6);

-- name: DeleteYearSchool :exec
UPDATE school_year SET deleted_at = $1 WHERE id = $2;

-- name: UpdateSchoolYear :exec
UPDATE school_year SET year = $1, start_at = $2, end_at = $3, updated_at = $4 WHERE id = $5;

-- name: FindOneSchoolYear :one
SELECT id, year, start_at, end_at FROM school_year WHERE id = $1 AND deleted_at IS NULL;

-- name: FindByYear :one
SELECT id, year, start_at, end_at FROM school_year WHERE year = $1 LIMIT 1;