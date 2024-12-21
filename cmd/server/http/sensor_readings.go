package http

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	db "github.com/christoff-linde/pih-core-go/internal/database"
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

	sensorReadings, err := s.db.GetSensorReadingMinutes(c, db.GetSensorReadingMinutesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	nextOffset := 0
	if sensorReadings == nil {
		nextOffset = 0
	} else if limit <= len(sensorReadings) {
		nextOffset = len(sensorReadings)
	} else if limit >= len(sensorReadings) {
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
	limit64, _ := strconv.ParseInt(c.DefaultQuery("limit", "7"), 10, 32)
	offset64, _ := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 32)
	limit := int32(limit64)
	offset := int32(offset64)
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

func (s *Server) GetSensorReadingsMinutesByIdHandler(c *gin.Context) {
	sensorID64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil || sensorID64 > math.MaxInt32 || sensorID64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}
	sensor_id := int32(sensorID64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}

	limit64, err := strconv.ParseInt(c.DefaultQuery("limit", "60"), 10, 32)
	if err != nil || limit64 > math.MaxInt32 || limit64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid limit parameter", err)})
		return
	}
	limit := int32(limit64)

	offset64, err := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 32)
	if err != nil || offset64 > math.MaxInt32 || offset64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid offset parameter", err)})
		return
	}
	offset := int32(offset64)

	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingMinutesById(c, db.GetSensorReadingMinutesByIdParams{
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
	} else if int32(len(sensorReadings)) < offset {

		nextOffset = 0
	}
	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorReadingsHourlyByIdHandler(c *gin.Context) {
	sensorID64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil || sensorID64 > math.MaxInt32 || sensorID64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}
	sensor_id := int32(sensorID64)
	limit64, err := strconv.ParseInt(c.DefaultQuery("limit", "24"), 10, 32)
	if err != nil || limit64 > math.MaxInt32 || limit64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid limit parameter", err)})
		return
	}
	limit := int32(limit64)
	offset64, err := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 32)
	if err != nil || offset64 > math.MaxInt32 || offset64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid offset parameter", err)})
		return
	}
	offset := int32(offset64)
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingHourlyById(c, db.GetSensorReadingHourlyByIdParams{
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
	} else if int32(len(sensorReadings)) < offset {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorReadingsDailyByIdHandler(c *gin.Context) {
	sensorID64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil || sensorID64 > math.MaxInt32 || sensorID64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}
	sensor_id := int32(sensorID64)
	limit64, err := strconv.ParseInt(c.DefaultQuery("limit", "7"), 10, 32)
	if err != nil || limit64 > math.MaxInt32 || limit64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid limit parameter", err)})
		return
	}
	limit := int32(limit64)
	offset64, err := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 32)
	if err != nil || offset64 > math.MaxInt32 || offset64 < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid offset parameter", err)})
		return
	}
	offset := int32(offset64)
	nextOffset := offset + limit

	sensorReadings, err := s.db.GetSensorReadingDailById(c, db.GetSensorReadingDailByIdParams{
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
	} else if int32(len(sensorReadings)) < offset {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_readings", len(sensorReadings)),
		"data":      sensorReadings,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}
