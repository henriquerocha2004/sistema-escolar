-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar

-- name: CreateStudent :exec
INSERT INTO students 
(id, first_name, last_name, birthday, rg_document, cpf_document, email, him_self_responsible, created_at, updated_at)
VALUES
($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);

-- name: FindByCPFDocument :one
SELECT id, first_name, last_name, birthday, rg_document, cpf_document, email, him_self_responsible FROM students WHERE cpf_document = $1 LIMIT 1;