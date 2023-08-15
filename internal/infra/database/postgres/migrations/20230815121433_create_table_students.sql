-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
-- +goose Up
-- +goose StatementBegin
CREATE TABLE students (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    rg_document VARCHAR(255),
    cpf_document VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    him_self_responsible BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE students;
-- +goose StatementEnd
