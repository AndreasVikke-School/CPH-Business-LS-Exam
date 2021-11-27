package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "redis_service/attendancecode"
	eh "redis_service/errorhandler"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAttendanceCodeProtoServer
}

func (s *server) CreateAttendanceCode(ctx context.Context, in *pb.AttendanceCodeCreate) (*pb.AttendanceCode, error) {
	code, unix, err := CreateAttendanceCodeInRedis(in.GetMinutesToLive())
	eh.PanicOnError(err, "Error when adding code to redis")
	return &pb.AttendanceCode{Code: code, Unix: unix}, nil
}

func (s *server) GetAttendanceCode(ctx context.Context, in *wrapperspb.Int64Value) (*pb.AttendanceCode, error) {
	code, unix, err := GetAttendanceCodeFromRedis(in.Value)
	eh.PanicOnError(err, "Error when getting code from redis")
	return &pb.AttendanceCode{Code: code, Unix: unix}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	eh.PanicOnError(err, "failed to listen")

	s := grpc.NewServer()
	pb.RegisterAttendanceCodeProtoServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	eh.PanicOnError(err, "failed to serve")
}
