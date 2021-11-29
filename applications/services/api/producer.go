package main

import (
	"context"
	"encoding/json"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type CheckIn struct {
	AttendanceCode  string `json:"attendanceCode"`
	StudentID       string `json:"studentId"`
	CurrentUnixTime int64  `json:"currentUnixTime"`
}

func Produce(checkIn CheckIn) {
	topic := "checkin"
	partition := 0

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := kafka.DialLeader(ctx, "tcp", configuration.Kafka.Service, topic, partition)
	eh.PanicOnError(err, "failed to dial leader")

	c, err := json.Marshal(checkIn)
	eh.PanicOnError(err, "Can't convert to JSON")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(c)})
	eh.PanicOnError(err, "failed to write messages")
}

func ProduceMessageToKafka(c *gin.Context) {
	var checkIn CheckIn
	// validate json from body against struct
	err := c.BindJSON(&checkIn)
	eh.PanicOnError(err, "Couldn't bind JSON")
	checkIn.CurrentUnixTime = time.Now().Unix()
	Produce(checkIn)
}
