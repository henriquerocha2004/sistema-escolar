-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
-- name: CreateParent :exec
INSERT INTO parents
(id, first_name, last_name, birthday, rg_document, cpf_document, student_id, email, created_at, updated_at)
VALUES
($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);