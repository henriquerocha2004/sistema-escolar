-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    capacity INTEGER NOT NULL,
    created_at TIMESTAMP(0) NULL,
    updated_at TIMESTAMP(0) NULL,
    deleted_at TIMESTAMP(0) NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
-- +goose StatementEnd
