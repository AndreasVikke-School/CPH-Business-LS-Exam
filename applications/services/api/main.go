package main

import (
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	configuration Configuration
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/attendance_code/:code", GetAttendanceCodeById)
	router.POST("/api/attendance_code/", CreateAttendanceCode)

	router.GET("/api/checkins/create_test_data", ProduceCheckInTestData)
	router.GET("/api/checkin/id/:id", GetCheckInById)
	router.GET("/api/checkins/", GetAllCheckIns)
	router.GET("/api/checkins/att/:attCode", GetCheckInsByAttendenceCode)
	router.GET("/api/checkins/student/:studentId", GetCheckInsByStudentId)
	router.GET("/api/checkins/time/:from/:to", GetCheckInsByTime)
	router.POST("/api/checkin/", ProduceCheckIn)

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	router.Run("0.0.0.0:8081")
}
