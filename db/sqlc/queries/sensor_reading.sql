-- name: CreateSensorReading :one
INSERT INTO sensor_readings (reading_timestamp, sensor_id, temperature, humidity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSensorReading :many
SELECT *
FROM sensor_readings
WHERE sensor_id = $1
LIMIT $2 OFFSET $3;

-- name: GetSensorReadingMinutes :many
SELECT *
FROM sensor_readings_minutes
LIMIT $1 OFFSET $2;

-- name: GetSensorReadingHourly :many
SELECT *
FROM sensor_readings_hourly
LIMIT $1 OFFSET $2;

-- name: GetSensorReadingDaily :many
SELECT *
FROM sensor_readings_daily
LIMIT $1 OFFSET $2;



---- name: SelectSensorReadings :many
-- SELECT time_bucket($1, time) as time_interval,
--   COUNT(*),
--   MAX(temperature) as max_temp,
--   AVG(temperature) as avg_temp,
--   MIN(temperature) as min_temp,
--   MAX(humidity) as max_hum,
--   AVG(humidity) as avg_hum,
--   MIN(humidity) as min_hum
-- FROM sensor_readings
-- WHERE time > NOW() - interval $2
-- GROUP BY time_interval, sensor_id
-- ORDER BY time_interval DESC, max_temp DESC;

-- -- name: CreateSensor :one
-- INSERT INTO sensors (sensor_name, created_at, updated_at)
-- VALUES ($1, $2, $3)
-- RETURNING *;
--
-- -- name: GetSensors :many
-- SELECT * FROM sensors LIMIT $1 OFFSET $2;
--
-- -- name: GetSensorById :one
-- SELECT * FROM sensors WHERE id=$1;
--
-- -- name: GetSensorByName :one
-- SELECT * FROM sensors WHERE sensor_name=$1;
--
-- -- name: UpdateSensor :execresult
-- UPDATE sensors
-- SET sensor_name = $2, updated_at = now()
-- WHERE id = $1
-- RETURNING *;
--
-- -- name: DeleteSensor :execresult
-- DELETE FROM sensors WHERE id=$1;
