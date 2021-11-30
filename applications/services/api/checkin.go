package main

import (
	"net/http"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	"github.com/gin-gonic/gin"
)

type CheckIn struct {
	AttendanceCode  string `json:"attendanceCode"`
	StudentID       string `json:"studentId"`
	CurrentUnixTime int64  `json:"currentUnixTime"`
}

func ProduceCheckIn(c *gin.Context) {
	var checkIn CheckIn
	// validate json from body against struct
	err := c.BindJSON(&checkIn)
	eh.PanicOnError(err, "Couldn't bind JSON")
	checkIn.CurrentUnixTime = time.Now().UnixNano() / 1000000
	ProduceCheckInToKafka(checkIn)
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
