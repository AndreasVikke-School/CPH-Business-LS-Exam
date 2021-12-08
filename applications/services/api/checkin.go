package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	eh "github.com/AndreasVikke-School/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	pb "github.com/AndreasVikke-School/CPH-Bussines-LS-Exam/applications/services/api/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type CheckIn struct {
	AttendanceCode  int64   `json:"attendanceCode"`
	StudentID       string  `json:"studentId"`
	CurrentUnixTime int64   `json:"currentUnixTime"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"long"`
}

func GetCheckInServiceClient() pb.CheckInServiceClient {
	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")

	return pb.NewCheckInServiceClient(conn)
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

func GetCheckInById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	eh.PanicOnError(err, "Id is not an int")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := GetCheckInServiceClient().GetCheckInById(ctx, &wrapperspb.Int64Value{Value: id})
	eh.PanicOnError(err, "Failed to get checkin")

	c.IndentedJSON(http.StatusOK, checkIn)
}

func GetCheckInsByStudentId(c *gin.Context) {
	studentId := c.Param("studentId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := GetCheckInServiceClient().GetCheckInsByStudentId(ctx, &wrapperspb.StringValue{Value: studentId})
	eh.PanicOnError(err, "Failed to get checkins")

	c.IndentedJSON(http.StatusOK, checkIn)
}

func GetCheckInsByAttendenceCode(c *gin.Context) {
	att := c.Param("attCode")
	attCode, err := strconv.ParseInt(att, 10, 64)
	eh.PanicOnError(err, "Attendance code is not an int")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := GetCheckInServiceClient().GetCheckInsByAttendenceCode(ctx, &wrapperspb.Int64Value{Value: attCode})
	eh.PanicOnError(err, "Failed to get checkins")

	c.IndentedJSON(http.StatusOK, checkIn)
}

func GetCheckInsByTime(c *gin.Context) {
	fromStr := c.Param("from")
	toStr := c.Param("to")
	from, err := strconv.ParseInt(fromStr, 10, 64)
	eh.PanicOnError(err, "Attendance code is not an int")
	to, err := strconv.ParseInt(toStr, 10, 64)
	eh.PanicOnError(err, "Attendance code is not an int")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := GetCheckInServiceClient().GetCheckInsByTime(ctx, &pb.TimeInterval{FromTime: from, ToTime: to})
	eh.PanicOnError(err, "Failed to get checkins")

	c.IndentedJSON(http.StatusOK, checkIn)
}

func GetAllCheckIns(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := GetCheckInServiceClient().GetAllCheckIns(ctx, &emptypb.Empty{})
	eh.PanicOnError(err, "Failed to get checkins")

	c.IndentedJSON(http.StatusOK, checkIn)
}
