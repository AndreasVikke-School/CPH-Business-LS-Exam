package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "api/rpc"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// getAlbums responds with the list of all albums as JSON.
func getPhotos(c *gin.Context) {
	conn, err := grpc.Dial("service_photos.test")
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

func main() {
	router := gin.Default()
	router.GET("/photos", getPhotos)

	router.Run("0.0.0.0:8080")
}
