package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/rpc"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	configuration Configuration
)

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
	router.Use(cors.Default())

	router.GET("/api/attendance_code/:code", GetAttendanceCodeById)
	router.POST("/api/attendance_code/", CreateAttendanceCode)
	router.GET("/api/checkin/:id", GetCheckInById)
	router.POST("/api/checkin/", ProduceCheckIn)

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	router.Run("0.0.0.0:8081")
}
