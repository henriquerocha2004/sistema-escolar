-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
-- name: CreatePhone :exec
INSERT INTO phones
(id, description, phone, owner_id)
VALUES
($1,$2,$3,$4);