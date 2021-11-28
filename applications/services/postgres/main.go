package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	eh "redis_service/errorhandler"
	pb "redis_service/rpc"

	"google.golang.org/grpc"
)

var (
	configuration Configuration
	port          = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCheckInProtoServer
}

// func (s *server) GetCheckInById(ctx context.Context, in *wrapperspb.Int64Value) (*pb.CheckIn, error) {
// 	// code, unix, err := GetAttendanceCodeFromRedis(in.Value, configuration)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return &pb.CheckIn{Code: code, Unix: unix}, nil
// }

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
