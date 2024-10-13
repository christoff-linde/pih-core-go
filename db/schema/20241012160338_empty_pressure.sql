-- +goose Up
-- +goose StatementBegin
ALTER TABLE sensor_readings
ALTER COLUMN pressure DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sensor_readings
ALTER COLUMN pressure SET NULL;
-- +goose StatementEnd
