package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent"
	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent/checkin"
	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/rpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	configuration Configuration
	port          = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCheckInServiceServer
}

func (s *server) GetCheckInById(ctx context.Context, in *wrapperspb.Int64Value) (*pb.CheckIn, error) {
	checkIn, err := GetCheckInByIdFromPostgres(context.Background(), in.Value, configuration)
	if err != nil {
		return nil, err
	}
	return &pb.CheckIn{Id: int64(checkIn.ID), AttendanceCode: checkIn.AttendanceCode, StudentId: checkIn.StudentId}, nil
}

func (s *server) GetAllCheckIns(ctx context.Context, empt *emptypb.Empty) (*pb.CheckIns, error) {
	checkIns, err := GetCheckIns(ctx, configuration)
	c, e := GetGrpcCheckIns(checkIns, err)
	return c, e
}

func (s *server) GetCheckInsByAttendenceCode(ctx context.Context, in *wrapperspb.Int64Value) (*pb.CheckIns, error) {
	checkIns, err := GetByAttendanceCode(ctx, in.Value, configuration)
	return GetGrpcCheckIns(checkIns, err)
}

func (s *server) GetCheckInsByStudentId(ctx context.Context, in *wrapperspb.StringValue) (*pb.CheckIns, error) {
	checkIns, err := GetByStudentId(ctx, in.Value, configuration)
	return GetGrpcCheckIns(checkIns, err)
}

func (s *server) GetCheckInsByTime(ctx context.Context, in *pb.TimeInterval) (*pb.CheckIns, error) {
	checkins, err := GetByTime(ctx, in.FromTime, in.ToTime, configuration)
	return GetGrpcCheckIns(checkins, err)
}

func GetGrpcCheckIns(checkIns []*ent.CheckIn, err error) (*pb.CheckIns, error) {
	if err != nil {
		return nil, err
	}

	var result []*pb.CheckIn
	for _, ci := range checkIns {
		result = append(result, &pb.CheckIn{Id: int64(ci.ID), AttendanceCode: ci.AttendanceCode, StudentId: ci.StudentId, CheckinTime: ci.CheckinTime, Status: MapStatusToValidity(ci.Status)})
	}

	return &pb.CheckIns{CheckIn: result}, nil
}

func (s *server) InsertCheckIn(ctx context.Context, in *pb.CheckIn) (*emptypb.Empty, error) {
	ci := &ent.CheckIn{AttendanceCode: in.AttendanceCode, StudentId: in.StudentId, Status: MapValidityToStatus(in.Status), CheckinTime: in.CheckinTime}
	_, err := InsertNewCheckin(ctx, ci, configuration)

	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func MapValidityToStatus(validity pb.Validity) checkin.Status {
	switch validity {
	case pb.Validity_SUCCESS:
		return checkin.StatusSuccess
	case pb.Validity_OUT_OF_TIME:
		return checkin.StatusOutOfTime
	case pb.Validity_NOT_FOUND:
		return checkin.StatusNotFound
	case pb.Validity_OUT_OF_RANGE:
		return checkin.StatusOutOfRange
	default:
		return checkin.StatusError
	}
}

func MapStatusToValidity(status checkin.Status) pb.Validity {
	switch status {
	case checkin.StatusSuccess:
		return pb.Validity_SUCCESS
	case checkin.StatusOutOfTime:
		return pb.Validity_OUT_OF_TIME
	case checkin.StatusNotFound:
		return pb.Validity_NOT_FOUND
	case checkin.StatusOutOfRange:
		return pb.Validity_OUT_OF_RANGE
	default:
		return pb.Validity_ERROR
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	eh.PanicOnError(err, "failed to listen")

	s := grpc.NewServer()
	pb.RegisterCheckInServiceServer(s, &server{})

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
