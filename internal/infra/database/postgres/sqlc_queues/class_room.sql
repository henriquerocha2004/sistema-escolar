-- name: CreateClass :exec
INSERT INTO class_room (id, status, identification, vacancies, vacancies_occupied, shift, level, localization, open_date, school_year_id, room_id, schedule_id, created_at, updated_at, type)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);

-- name: UpdateClass :exec
UPDATE class_room SET
        status = $1,
        identification = $2,
        vacancies = $3,
        vacancies_occupied = $4,
        shift = $5,
        level = $6,
        localization = $7,
        open_date = $8,
        school_year_id = $9,
        room_id = $10,
        schedule_id = $11,
        updated_at = $12
WHERE id = $13;

-- name: DeleteClass :exec
UPDATE class_room SET deleted_at = $1 WHERE id = $2;

-- name: FindClassById :one
SELECT id,
       status,
       identification,
       vacancies,
       vacancies_occupied,
       shift,
       level,
       localization,
       open_date,
       school_year_id,
       room_id,
      schedule_id,
      type
FROM class_room
    WHERE id = $1
        AND deleted_at IS NULL;

-- name: FindClassByIdLock :one

SELECT id,
       status,
       identification,
       vacancies,
       vacancies_occupied,
       shift,
       level,
       localization,
       open_date,
       school_year_id,
       room_id,
      schedule_id,
      type
FROM class_room
    WHERE id = $1
        AND deleted_at IS NULL
        FOR UPDATE;


-- name: UpdateVacancyOccupied :exec
UPDATE class_room 
    SET vacancies_occupied = $1, 
        updated_at = $2 
WHERE 
    id = $3 
    AND vacancies_occupied < vacancies 
    AND deleted_at IS NULL 
    RETURNING vacancies_occupied;