-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS sensor_readings_minutes
WITH (timescaledb.continuous) AS
SELECT 
  time_bucket('1 minute', sensor_readings.reading_timestamp) AS minute,
  sensor_id,
  min(temperature) as min_temperature,
  avg(temperature) as avg_temperature,
  max(temperature) as max_temperature,
  min(humidity) as min_humidity,
  avg(humidity) as avg_humidity,
  max(humidity) as max_humidity
FROM sensor_readings
GROUP BY minute, sensor_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP MATERIALIZED VIEW IF EXISTS sensor_readings_minutes;
-- +goose StatementEnd