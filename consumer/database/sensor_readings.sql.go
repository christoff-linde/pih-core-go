// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sensor_readings.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSensorReading = `-- name: CreateSensorReading :one
INSERT INTO sensor_readings (reading_timestamp, sensor_id, temperature, humidity, pressure)
VALUES ($1, $2, $3, $4, $5)
RETURNING reading_timestamp, sensor_id, temperature, humidity, pressure
`

type CreateSensorReadingParams struct {
	ReadingTimestamp pgtype.Timestamptz `json:"reading_timestamp"`
	SensorID         pgtype.Int4        `json:"sensor_id"`
	Temperature      pgtype.Float8      `json:"temperature"`
	Humidity         pgtype.Float8      `json:"humidity"`
	Pressure         pgtype.Float8      `json:"pressure"`
}

func (q *Queries) CreateSensorReading(ctx context.Context, arg CreateSensorReadingParams) (SensorReading, error) {
	row := q.db.QueryRow(ctx, createSensorReading,
		arg.ReadingTimestamp,
		arg.SensorID,
		arg.Temperature,
		arg.Humidity,
		arg.Pressure,
	)
	var i SensorReading
	err := row.Scan(
		&i.ReadingTimestamp,
		&i.SensorID,
		&i.Temperature,
		&i.Humidity,
		&i.Pressure,
	)
	return i, err
}
