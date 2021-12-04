package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProduceCheckInTestData(c *gin.Context) {
	var checkins []CheckIn

	for a := 0; a < 10; a++ {
		attendancecode := CreateAttendanceCodeInRedis(CreateAttendanceCodeType{
			MinutesToLive: 1,
			Lat:           15,
			Long:          15,
		})
		for a := 0; a < rand.Intn(35)+10; a++ {
			checkins = append(checkins, CheckIn{
				AttendanceCode:  attendancecode.GetCode(),
				StudentID:       "azv@efio.dk",
				CurrentUnixTime: attendancecode.GetUnix(),
				Latitude:        attendancecode.GetLat(),
				Longitude:       attendancecode.GetLong(),
			})
		}
		for a := 0; a < rand.Intn(2); a++ {
			checkins = append(checkins, CheckIn{
				AttendanceCode:  attendancecode.GetCode(),
				StudentID:       "azv@efio.dk",
				CurrentUnixTime: 999999999999999999,
				Latitude:        attendancecode.GetLat(),
				Longitude:       attendancecode.GetLong(),
			})
		}
	}

	for _, checkin := range checkins {
		ProduceCheckInToKafka(checkin)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
