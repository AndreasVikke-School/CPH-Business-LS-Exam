package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	eh "postgres_service/errorhandler"
	pb "postgres_service/rpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	configuration Configuration
	port          = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCheckInProtoServer
}

func (s *server) GetCheckInById(ctx context.Context, in *wrapperspb.Int64Value) (*pb.CheckIn, error) {
	checkIn, err := GetCheckInByIdFromPostgres(context.Background(), in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.CheckIn{Id: int64(checkIn.ID), AttendanceCode: checkIn.AttendanceCode, StudentId: checkIn.StudentId}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	eh.PanicOnError(err, "failed to listen")

	s := grpc.NewServer()
	pb.RegisterCheckInProtoServer(s, &server{})

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	MigratePostgres(configuration)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	eh.PanicOnError(err, "failed to serve")
}
