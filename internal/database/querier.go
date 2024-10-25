// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	//CreateSensor
	//
	//  INSERT INTO sensors (sensor_name, created_at, updated_at)
	//  VALUES ($1, $2, $3)
	//  RETURNING id, sensor_name, created_at, updated_at
	CreateSensor(ctx context.Context, arg CreateSensorParams) (Sensor, error)
	//CreateSensorMetadata
	//
	//  INSERT INTO sensor_metadata ( id, sensor_id, sensor_type, manufacturer, model_number, sensor_location, additional_data )
	//  VALUES ($1, $2, $3, $4, $5, $6, $7)
	//  RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
	CreateSensorMetadata(ctx context.Context, arg CreateSensorMetadataParams) (SensorMetadatum, error)
	//CreateSensorReading
	//
	//  INSERT INTO sensor_readings (sensor_id, temperature, humidity)
	//  VALUES ($1, $2, $3)
	//  RETURNING reading_timestamp, sensor_id, temperature, humidity, pressure
	CreateSensorReading(ctx context.Context, arg CreateSensorReadingParams) (SensorReading, error)
	//DeleteSensor
	//
	//  DELETE FROM sensors WHERE id=$1
	DeleteSensor(ctx context.Context, id int32) (pgconn.CommandTag, error)
	//DeleteSensorMetadata
	//
	//  DELETE FROM sensor_metadata WHERE id = $1
	DeleteSensorMetadata(ctx context.Context, id int32) (pgconn.CommandTag, error)
	//GetSensorById
	//
	//  SELECT id, sensor_name, created_at, updated_at FROM sensors WHERE id=$1
	GetSensorById(ctx context.Context, id int32) (Sensor, error)
	//GetSensorByName
	//
	//  SELECT id, sensor_name, created_at, updated_at FROM sensors WHERE sensor_name=$1
	GetSensorByName(ctx context.Context, sensorName string) (Sensor, error)
	//GetSensorMetadata
	//
	//  SELECT id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
	//  FROM sensor_metadata
	//  LIMIT $1 OFFSET $2
	GetSensorMetadata(ctx context.Context, arg GetSensorMetadataParams) ([]SensorMetadatum, error)
	//GetSensorMetadataForSensorId
	//
	//  SELECT sensors.id, sensors.sensor_name, sensors.created_at, sensors.updated_at
	//  FROM sensors
	//           JOIN sensor_metadata ON sensors.id = sensor_metadata.sensor_id
	//  WHERE sensor_metadata.sensor_id = $1
	GetSensorMetadataForSensorId(ctx context.Context, sensorID int32) (Sensor, error)
	//GetSensorReading
	//
	//  SELECT reading_timestamp, sensor_id, temperature, humidity, pressure
	//  FROM sensor_readings
	//  WHERE sensor_id = $1
	//  LIMIT $2 OFFSET $3
	GetSensorReading(ctx context.Context, arg GetSensorReadingParams) ([]SensorReading, error)
	//GetSensorReadingDaily
	//
	//  SELECT day, sensor_id, min_temperature, avg_temperature, max_temperature, min_humidity, avg_humidity, max_humidity
	//  FROM sensor_readings_daily
	//  LIMIT $1 OFFSET $2
	GetSensorReadingDaily(ctx context.Context, arg GetSensorReadingDailyParams) ([]SensorReadingsDaily, error)
	//GetSensorReadingHourly
	//
	//  SELECT hour, sensor_id, min_temperature, avg_temperature, max_temperature, min_humidity, avg_humidity, max_humidity
	//  FROM sensor_readings_hourly
	//  LIMIT $1 OFFSET $2
	GetSensorReadingHourly(ctx context.Context, arg GetSensorReadingHourlyParams) ([]SensorReadingsHourly, error)
	//GetSensorReadingMinutes
	//
	//  SELECT minute, sensor_id, min_temperature, avg_temperature, max_temperature, min_humidity, avg_humidity, max_humidity
	//  FROM sensor_readings_minutes
	//  LIMIT $1 OFFSET $2
	GetSensorReadingMinutes(ctx context.Context, arg GetSensorReadingMinutesParams) ([]SensorReadingsMinute, error)
	//GetSensors
	//
	//  SELECT id, sensor_name, created_at, updated_at FROM sensors LIMIT $1 OFFSET $2
	GetSensors(ctx context.Context, arg GetSensorsParams) ([]Sensor, error)
	//UpdateSensor
	//
	//  UPDATE sensors
	//  SET sensor_name = $2, updated_at = now()
	//  WHERE id = $1
	//  RETURNING id, sensor_name, created_at, updated_at
	UpdateSensor(ctx context.Context, arg UpdateSensorParams) (pgconn.CommandTag, error)
	//UpdateSensorMetadata
	//
	//  UPDATE sensor_metadata
	//  SET sensor_id = $1, sensor_type = $3, manufacturer = $4, model_number = $5, sensor_location = $6, additional_data = $7
	//  WHERE id = $2
	//  RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
	UpdateSensorMetadata(ctx context.Context, arg UpdateSensorMetadataParams) (pgconn.CommandTag, error)
}

var _ Querier = (*Queries)(nil)
