// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sensor.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSensor = `-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, sensor_unique_id, sensor_location, sensor_type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, sensor_name, sensor_unique_id, sensor_location, sensor_type, created_at, updated_at
`

type CreateSensorParams struct {
	SensorName     string             `json:"sensor_name"`
	SensorUniqueID pgtype.UUID        `json:"sensor_unique_id"`
	SensorLocation pgtype.Text        `json:"sensor_location"`
	SensorType     pgtype.Text        `json:"sensor_type"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateSensor(ctx context.Context, arg CreateSensorParams) (Sensor, error) {
	row := q.db.QueryRow(ctx, createSensor,
		arg.SensorName,
		arg.SensorUniqueID,
		arg.SensorLocation,
		arg.SensorType,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Sensor
	err := row.Scan(
		&i.ID,
		&i.SensorName,
		&i.SensorUniqueID,
		&i.SensorLocation,
		&i.SensorType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
