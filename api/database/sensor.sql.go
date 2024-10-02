// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sensor.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSensor = `-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, created_at, updated_at)
VALUES ($1, $2, $3)
RETURNING id, sensor_name, created_at, updated_at
`

type CreateSensorParams struct {
	SensorName string             `json:"sensor_name"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateSensor(ctx context.Context, arg CreateSensorParams) (Sensor, error) {
	row := q.db.QueryRow(ctx, createSensor, arg.SensorName, arg.CreatedAt, arg.UpdatedAt)
	var i Sensor
	err := row.Scan(
		&i.ID,
		&i.SensorName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSensor = `-- name: DeleteSensor :execresult
DELETE FROM sensors WHERE id=$1
`

func (q *Queries) DeleteSensor(ctx context.Context, id int32) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteSensor, id)
}

const getSensorById = `-- name: GetSensorById :one
SELECT id, sensor_name, created_at, updated_at FROM sensors WHERE id=$1
`

func (q *Queries) GetSensorById(ctx context.Context, id int32) (Sensor, error) {
	row := q.db.QueryRow(ctx, getSensorById, id)
	var i Sensor
	err := row.Scan(
		&i.ID,
		&i.SensorName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSensorByName = `-- name: GetSensorByName :one
SELECT id, sensor_name, created_at, updated_at FROM sensors WHERE sensor_name=$1
`

func (q *Queries) GetSensorByName(ctx context.Context, sensorName string) (Sensor, error) {
	row := q.db.QueryRow(ctx, getSensorByName, sensorName)
	var i Sensor
	err := row.Scan(
		&i.ID,
		&i.SensorName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSensors = `-- name: GetSensors :many
SELECT id, sensor_name, created_at, updated_at FROM sensors LIMIT $1 OFFSET $2
`

type GetSensorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetSensors(ctx context.Context, arg GetSensorsParams) ([]Sensor, error) {
	rows, err := q.db.Query(ctx, getSensors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sensor
	for rows.Next() {
		var i Sensor
		if err := rows.Scan(
			&i.ID,
			&i.SensorName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSensor = `-- name: UpdateSensor :execresult
UPDATE sensors
SET sensor_name = $2, updated_at = now()
WHERE id = $1
RETURNING id, sensor_name, created_at, updated_at
`

type UpdateSensorParams struct {
	ID         int32  `json:"id"`
	SensorName string `json:"sensor_name"`
}

func (q *Queries) UpdateSensor(ctx context.Context, arg UpdateSensorParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, updateSensor, arg.ID, arg.SensorName)
}
