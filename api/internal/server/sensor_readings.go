package server

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/christoff-linde/pih-core-go/api/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) CreateSensorReadingHandler(c *gin.Context) {
	var sensorReading db.SensorReading

	if err := c.BindJSON(&sensorReading); err != nil {
		resp := gin.H{
			"message": "Invalid input",
			"error":   err,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	sensorReading, err := s.db.CreateSensorReading(c, db.CreateSensorReadingParams{
		SensorID:    sensorReading.SensorID,
		Temperature: sensorReading.Temperature,
		Humidity:    sensorReading.Humidity,
	})
	if err != nil {
		resp := gin.H{
			"message": "Failed to create sensor reading.",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Created sensor_reading %v", sensorReading),
		"data":    sensorReading,
	}
	c.JSON(http.StatusCreated, resp)
}

func (s *Server) GetSensorReadingsHandler(c *gin.Context) {
	sensor_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "60"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReading(c, db.GetSensorReadingParams{
		SensorID: pgtype.Int4{
			Int32: int32(sensor_id),
			Valid: true,
		},
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensorReadings == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorReadingsMinutesHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "60"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingMinutes(c, db.GetSensorReadingMinutesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensorReadings == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorReadingsHourlyHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "24"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingHourly(c, db.GetSensorReadingHourlyParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensorReadings == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorReadingsDailyHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "7"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingDaily(c, db.GetSensorReadingDailyParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensorReadings == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}
