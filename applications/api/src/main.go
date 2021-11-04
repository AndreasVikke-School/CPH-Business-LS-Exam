package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	pb "api/rpc"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// getAlbums responds with the list of all albums as JSON.
func getPhotos(c *gin.Context) {
	conn, err := grpc.Dial("service-photos.test:50051")
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPhotoProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	photos, err := client.GetAllPhotos(ctx, &empty.Empty{})

	c.IndentedJSON(http.StatusOK, photos)
}

func getPhotoById(c *gin.Context) {
	conn, err := grpc.Dial("service-photos.test:50051")
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPhotoProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	photo, err := client.GetPhotoById(ctx, wrapperspb.Int32(int32(id)))

	c.IndentedJSON(http.StatusOK, photo)
}

func main() {
	router := gin.Default()
	router.GET("/photos", getPhotos)
	router.GET("/photos/:id", getPhotos)

	router.Run("0.0.0.0:8080")
}
