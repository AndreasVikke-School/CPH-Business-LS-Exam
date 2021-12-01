package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func CreateAttendanceCode(c *gin.Context) {
	minutesToLiveStr := c.Param("minutesToLive")
	minutesToLive, err := strconv.ParseInt(minutesToLiveStr, 10, 64)
	eh.PanicOnError(err, "Minutes to live is not an int")

	conn, err := grpc.Dial(configuration.Redis.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewAttendanceCodeProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newAttendanceCode := &pb.AttendanceCodeCreate{MinutesToLive: minutesToLive}
	attendancecode, err := client.CreateAttendanceCode(ctx, newAttendanceCode)
	eh.PanicOnError(err, "Failed to create attendance code")

	c.IndentedJSON(http.StatusOK, attendancecode)
}

func GetAttendanceCodeById(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.ParseInt(codeStr, 10, 64)
	eh.PanicOnError(err, "Minutes to live is not an int")

	conn, err := grpc.Dial(configuration.Redis.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewAttendanceCodeProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	attendancecode, err := client.GetAttendanceCodeById(ctx, &wrapperspb.Int64Value{Value: code})
	eh.PanicOnError(err, "Failed to get attendance code")

	if attendancecode.GetCode() >= 0 {
		c.IndentedJSON(http.StatusOK, attendancecode)
	} else {
		c.Status(http.StatusOK)
	}
}
