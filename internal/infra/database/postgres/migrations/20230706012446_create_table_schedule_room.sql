-- +goose Up
-- +goose StatementBegin
CREATE TABLE room_schedule (
    id BIGINT PRIMARY KEY,
    room_id UUID NOT NULL,
    schedule_id UUID NOT NULL,
    school_year_id UUID NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE room_schedule ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms (id);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE room_schedule ADD CONSTRAINT fk_schedule FOREIGN KEY (schedule_id) REFERENCES class_schedule (id);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE room_schedule ADD CONSTRAINT fk_school_year FOREIGN KEY (school_year_id) REFERENCES school_year (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE room_schedule;
-- +goose StatementEnd
