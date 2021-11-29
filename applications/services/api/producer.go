package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	topic := "my-topic"
	partition := 1
	fmt.Println("hej vikke")
	conn, err := kafka.DialLeader(context.Background(), "tcp", configuration.Kafka.Service, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	c, err := json.Marshal(checkIn)
	eh.PanicOnError(err, "Can't convert to JSON")
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(c)})
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func ProduceMessageToKafka(c *gin.Context) {
	var checkIn CheckIn
	// validate json from body against struct
	err := c.BindJSON(&checkIn)
	eh.PanicOnError(err, "Couldn't bind JSON")
	checkIn.CurrentUnixTime = time.Now().Unix()
	Produce(checkIn)
}
