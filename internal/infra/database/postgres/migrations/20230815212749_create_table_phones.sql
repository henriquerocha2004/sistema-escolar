-- Active: 1691937846246@@127.0.0.1@9500@sistema-escolar
-- +goose Up
-- +goose StatementBegin
CREATE TABLE phones (
    id UUID PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE phones;
-- +goose StatementEnd
