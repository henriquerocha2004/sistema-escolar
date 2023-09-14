-- Active: 1691937846246@@127.0.0.1@9500@postgres

-- name: CreateRegistration :exec
INSERT INTO registrations 
    (id, code, class_room_id, shift, student_id, 
     service_id, monthly_fee, installments_quantity,
     enrollment_fee, due_date, month_duration, status,
     enrollment_date, school_year_id, created_at, updated_at)
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16);