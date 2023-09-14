-- +goose Up
-- +goose StatementBegin
CREATE TABLE registrations (
    id UUID PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    class_room_id UUID,
    shift VARCHAR(255),
    student_id UUID NOT NULL,
    service_id UUID NOT NULL,
    monthly_fee NUMERIC(10,2) NOT NULL,
    installments_quantity INTEGER NOT NULL,
    enrollment_fee NUMERIC(10,2),
    due_date TIMESTAMP,
    month_duration INTEGER,
    status VARCHAR(255) NOT NULL,
    enrollment_date TIMESTAMP,
    school_year_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE registrations;
-- +goose StatementEnd
