package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	configuration Configuration
)

// getAlbums responds with the list of all albums as JSON.
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

	c.IndentedJSON(http.StatusOK, attendancecode)
}

func GetCheckInById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	eh.PanicOnError(err, "Id is not an int")

	conn, err := grpc.Dial(configuration.Postgres.Service, grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewCheckInProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkIn, err := client.GetCheckInById(ctx, &wrapperspb.Int64Value{Value: id})
	eh.PanicOnError(err, "Failed to get attendance code")

	c.IndentedJSON(http.StatusOK, checkIn)
}

func main() {
	router := gin.Default()
	router.GET("/api/attendance_code/:code", GetAttendanceCodeById)
	router.POST("/api/attendance_code/:minutesToLive", CreateAttendanceCode)
	router.GET("/api/checkin/:id", GetCheckInById)

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	router.Run("0.0.0.0:8080")
}
