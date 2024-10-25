package subscriber

import (
	"context"

	db "github.com/christoff-linde/pih-core-go/internal/database"
)

func (appConfig *AppConfig) handleCreateSensor() (db.Sensor, error) {
	sensor, err := appConfig.DB.CreateSensor(context.Background(), db.CreateSensorParams{
		SensorName: "esp32-test-02",
	})

	return sensor, err
}

func (appConfig *AppConfig) handleGetSensorBySensorId(id int32) db.Sensor {
	sensor, err := appConfig.DB.GetSensorById(context.Background(), id)
	failOnError(err, "Could not fetch sensor")

	return sensor
}

func (appConfig *AppConfig) handleGetSensorByName(name string) (db.Sensor, error) {
	sensor, err := appConfig.DB.GetSensorByName(context.Background(), name)

	return sensor, err
}

func (appConfig *AppConfig) handleCreateSensorMetadata(sensor db.Sensor) db.SensorMetadatum {
	metadata, err := appConfig.DB.CreateSensorMetadata(context.Background(), db.CreateSensorMetadataParams{
		ID:       1,
		SensorID: sensor.ID,
	})
	failOnError(err, "Failed to create sensorMetadata")

	return metadata
}
