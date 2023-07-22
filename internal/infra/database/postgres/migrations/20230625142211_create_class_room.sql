-- +goose Up
-- +goose StatementBegin
CREATE TABLE class_room (
    id UUID PRIMARY KEY,
    active BOOLEAN NOT NULL DEFAULT false,
    status VARCHAR(255) NOT NULL,
    identification VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    vacancies INTEGER NOT NULL,
    vacancies_occupied INTEGER NOT NULL,
    shift VARCHAR(255) NOT NULL,
    level VARCHAR(255) NOT NULL,
    localization VARCHAR(255),
    open_date TIMESTAMP NOT NULL,
    school_year_id UUID NOT NULL,
    room_id UUID,
    schedule_id UUID NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE class_room;
-- +goose StatementEnd