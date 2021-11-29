package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/redis/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/redis/rpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	configuration Configuration
	port          = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAttendanceCodeProtoServer
}

func (s *server) CreateAttendanceCode(ctx context.Context, in *pb.AttendanceCodeCreate) (*pb.AttendanceCode, error) {
	code, unix, err := CreateAttendanceCodeInRedis(in.GetMinutesToLive(), configuration)
	// eh.PanicOnError(err, "Error when adding code to redis")
	if err != nil {
		return nil, err
	}
	return &pb.AttendanceCode{Code: code, Unix: unix}, nil
}

func (s *server) GetAttendanceCodeById(ctx context.Context, in *wrapperspb.Int64Value) (*pb.AttendanceCode, error) {
	code, unix, err := GetAttendanceCodeFromRedis(in.Value, configuration)
	// eh.PanicOnError(err, "Error when getting code from redis")
	if err != nil {
		return nil, err
	}
	return &pb.AttendanceCode{Code: code, Unix: unix}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	eh.PanicOnError(err, "failed to listen")

	s := grpc.NewServer()
	pb.RegisterAttendanceCodeProtoServer(s, &server{})

	if len(os.Args) >= 2 {
		configuration = getConfig(os.Args[1])
	} else {
		configuration = getConfig("dev")
	}

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	eh.PanicOnError(err, "failed to serve")
}
