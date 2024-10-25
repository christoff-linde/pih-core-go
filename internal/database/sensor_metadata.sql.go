// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sensor_metadata.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSensorMetadata = `-- name: CreateSensorMetadata :one
INSERT INTO sensor_metadata ( id, sensor_id, sensor_type, manufacturer, model_number, sensor_location, additional_data )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
`

type CreateSensorMetadataParams struct {
	ID             int32       `db:"id" json:"id"`
	SensorID       int32       `db:"sensor_id" json:"sensor_id"`
	SensorType     pgtype.Text `db:"sensor_type" json:"sensor_type"`
	Manufacturer   pgtype.Text `db:"manufacturer" json:"manufacturer"`
	ModelNumber    pgtype.Text `db:"model_number" json:"model_number"`
	SensorLocation pgtype.Text `db:"sensor_location" json:"sensor_location"`
	AdditionalData []byte      `db:"additional_data" json:"additional_data"`
}

// CreateSensorMetadata
//
//	INSERT INTO sensor_metadata ( id, sensor_id, sensor_type, manufacturer, model_number, sensor_location, additional_data )
//	VALUES ($1, $2, $3, $4, $5, $6, $7)
//	RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
func (q *Queries) CreateSensorMetadata(ctx context.Context, arg CreateSensorMetadataParams) (SensorMetadatum, error) {
	row := q.db.QueryRow(ctx, createSensorMetadata,
		arg.ID,
		arg.SensorID,
		arg.SensorType,
		arg.Manufacturer,
		arg.ModelNumber,
		arg.SensorLocation,
		arg.AdditionalData,
	)
	var i SensorMetadatum
	err := row.Scan(
		&i.ID,
		&i.SensorType,
		&i.Manufacturer,
		&i.ModelNumber,
		&i.SensorLocation,
		&i.InstallationTime,
		&i.UpdatedAt,
		&i.AdditionalData,
		&i.SensorID,
	)
	return i, err
}

const deleteSensorMetadata = `-- name: DeleteSensorMetadata :execresult
DELETE FROM sensor_metadata WHERE id = $1
`

// DeleteSensorMetadata
//
//	DELETE FROM sensor_metadata WHERE id = $1
func (q *Queries) DeleteSensorMetadata(ctx context.Context, id int32) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteSensorMetadata, id)
}

const getSensorMetadata = `-- name: GetSensorMetadata :many
SELECT id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
FROM sensor_metadata
LIMIT $1 OFFSET $2
`

type GetSensorMetadataParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

// GetSensorMetadata
//
//	SELECT id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
//	FROM sensor_metadata
//	LIMIT $1 OFFSET $2
func (q *Queries) GetSensorMetadata(ctx context.Context, arg GetSensorMetadataParams) ([]SensorMetadatum, error) {
	rows, err := q.db.Query(ctx, getSensorMetadata, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SensorMetadatum
	for rows.Next() {
		var i SensorMetadatum
		if err := rows.Scan(
			&i.ID,
			&i.SensorType,
			&i.Manufacturer,
			&i.ModelNumber,
			&i.SensorLocation,
			&i.InstallationTime,
			&i.UpdatedAt,
			&i.AdditionalData,
			&i.SensorID,
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

const getSensorMetadataForSensorId = `-- name: GetSensorMetadataForSensorId :one
SELECT sensors.id, sensors.sensor_name, sensors.created_at, sensors.updated_at
FROM sensors
         JOIN sensor_metadata ON sensors.id = sensor_metadata.sensor_id
WHERE sensor_metadata.sensor_id = $1
`

// GetSensorMetadataForSensorId
//
//	SELECT sensors.id, sensors.sensor_name, sensors.created_at, sensors.updated_at
//	FROM sensors
//	         JOIN sensor_metadata ON sensors.id = sensor_metadata.sensor_id
//	WHERE sensor_metadata.sensor_id = $1
func (q *Queries) GetSensorMetadataForSensorId(ctx context.Context, sensorID int32) (Sensor, error) {
	row := q.db.QueryRow(ctx, getSensorMetadataForSensorId, sensorID)
	var i Sensor
	err := row.Scan(
		&i.ID,
		&i.SensorName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSensorMetadata = `-- name: UpdateSensorMetadata :execresult
UPDATE sensor_metadata
SET sensor_id = $1, sensor_type = $3, manufacturer = $4, model_number = $5, sensor_location = $6, additional_data = $7
WHERE id = $2
RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
`

type UpdateSensorMetadataParams struct {
	SensorID       int32       `db:"sensor_id" json:"sensor_id"`
	ID             int32       `db:"id" json:"id"`
	SensorType     pgtype.Text `db:"sensor_type" json:"sensor_type"`
	Manufacturer   pgtype.Text `db:"manufacturer" json:"manufacturer"`
	ModelNumber    pgtype.Text `db:"model_number" json:"model_number"`
	SensorLocation pgtype.Text `db:"sensor_location" json:"sensor_location"`
	AdditionalData []byte      `db:"additional_data" json:"additional_data"`
}

// UpdateSensorMetadata
//
//	UPDATE sensor_metadata
//	SET sensor_id = $1, sensor_type = $3, manufacturer = $4, model_number = $5, sensor_location = $6, additional_data = $7
//	WHERE id = $2
//	RETURNING id, sensor_type, manufacturer, model_number, sensor_location, installation_time, updated_at, additional_data, sensor_id
func (q *Queries) UpdateSensorMetadata(ctx context.Context, arg UpdateSensorMetadataParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, updateSensorMetadata,
		arg.SensorID,
		arg.ID,
		arg.SensorType,
		arg.Manufacturer,
		arg.ModelNumber,
		arg.SensorLocation,
		arg.AdditionalData,
	)
}