package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/christoff-linde/pih-core-go/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AppConfig struct {
	DB *db.Queries
}

type DeviceData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func StartSubscriber(appCfg AppConfig, brokerUrl string) {
	conn, err := amqp.Dial(brokerUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclarePassive("iot", "topic", true, true, false, false, nil)
	failOnError(err, "IOT exchange does not exist")

	iotQueue, err := channel.QueueDeclare("iot", true, false, false, false, nil)
	failOnError(err, "Failed to declare iot queue")

	err = channel.QueueBind(iotQueue.Name, "pih", "iot", false, nil)
	failOnError(err, "Failed to bind iot queue")

	iotMsgs, err := channel.ConsumeWithContext(context.Background(), iotQueue.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range iotMsgs {
			log.Printf("Received a message from: %v: %v", iotQueue.Name, d.Body)

			var deviceData DeviceData
			err := json.Unmarshal([]byte(d.Body), &deviceData)
			if err != nil {
				log.Printf("Error parsing JSON: %v", err)
			}

			sensor, err := appCfg.DB.GetSensorByName(context.Background(), deviceData.DeviceID)
			if err != nil {
				log.Printf("Sensor %v not found: %v", deviceData.DeviceID, err)
			} else {
				if deviceData.Temperature > 0 || deviceData.Humidity > 0 {
					sensorReading, err := appCfg.DB.CreateSensorReading(context.Background(), db.CreateSensorReadingParams{
						SensorID: pgtype.Int4{
							Int32: sensor.ID,
							Valid: true,
						},
						Temperature: pgtype.Float8{
							Float64: deviceData.Temperature,
							Valid:   true,
						},
						Humidity: pgtype.Float8{
							Float64: deviceData.Humidity,
							Valid:   true,
						},
					})
					if err != nil {
						log.Printf("Failed to create sensorReading: %v", err)
					}
					fmt.Printf("Added: %v", sensorReading)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever
}
