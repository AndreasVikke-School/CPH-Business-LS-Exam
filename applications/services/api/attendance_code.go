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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type CreateAttendanceCodeType struct {
	MinutesToLive int64   `json:"minutesToLive"`
	Lat           float64 `json:"lat"`
	Long          float64 `json:"long"`
}

func CreateAttendanceCode(c *gin.Context) {
	var codeCreate CreateAttendanceCodeType
	err := c.BindJSON(&codeCreate)
	eh.PanicOnError(err, "Couldn't bind JSON")

	attendancecode := CreateAttendanceCodeInRedis(codeCreate)

	c.IndentedJSON(http.StatusOK, attendancecode)
}

func CreateAttendanceCodeInRedis(codeCreate CreateAttendanceCodeType) *pb.AttendanceCode {
	conn, err := grpc.Dial(configuration.Redis.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewAttendanceCodeProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newAttendanceCode := &pb.AttendanceCodeCreate{MinutesToLive: codeCreate.MinutesToLive, Lat: codeCreate.Lat, Long: codeCreate.Long}
	attendancecode, err := client.CreateAttendanceCode(ctx, newAttendanceCode)
	eh.PanicOnError(err, "Failed to create attendance code")
	return attendancecode
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
