-- +goose Up
-- +goose StatementBegin
CREATE TABLE school_year(
    id UUID PRIMARY KEY,
    year VARCHAR(255) NOT NULL UNIQUE,
    start_at DATE NOT NULL,
    end_at DATE NOT NULL,
    created_at TIMESTAMP(0) NULL,
    updated_at TIMESTAMP(0) NULL,
    deleted_at TIMESTAMP(0) NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE class_schedule ADD CONSTRAINT fk_schedule_school_year FOREIGN KEY (school_year_id) REFERENCES school_year (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE school_year;
-- +goose StatementEnd
