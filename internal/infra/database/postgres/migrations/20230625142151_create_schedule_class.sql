-- +goose Up
-- +goose StatementBegin
CREATE TABLE class_schedule (
    id UUID PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    schedule VARCHAR(255) NOT NULL,
    school_year_id UUID NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE class_schedule;
-- +goose StatementEnd