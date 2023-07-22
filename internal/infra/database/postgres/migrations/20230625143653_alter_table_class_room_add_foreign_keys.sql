-- +goose Up
-- +goose StatementBegin
ALTER TABLE class_room ADD CONSTRAINT fk_class_room_room FOREIGN KEY (room_id) REFERENCES rooms (id);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE class_room ADD CONSTRAINT fk_class_room_schedule FOREIGN KEY (schedule_id) REFERENCES class_schedule (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
