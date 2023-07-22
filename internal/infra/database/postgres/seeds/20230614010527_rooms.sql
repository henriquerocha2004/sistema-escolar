-- +goose Up
-- +goose StatementBegin
INSERT INTO rooms (id, code, description, capacity)
VALUES
    (UUID_TO_BIN(UUID()), '1S', 'SALA 1', 25),
    (UUID_TO_BIN(UUID()), '2S', 'SALA 2', 25),
    (UUID_TO_BIN(UUID()), '3S', 'SALA 3', 25),
    (UUID_TO_BIN(UUID()), '4S', 'SALA 4', 25),
    (UUID_TO_BIN(UUID()), '5S', 'SALA 5', 25),
    (UUID_TO_BIN(UUID()), '6S', 'SALA 6', 25);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
