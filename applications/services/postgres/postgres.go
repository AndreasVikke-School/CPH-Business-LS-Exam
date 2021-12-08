package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AndreasVikke-School/CPH-Bussines-LS-Exam/applications/services/postgres/ent"
	"github.com/AndreasVikke-School/CPH-Bussines-LS-Exam/applications/services/postgres/ent/checkin"

	_ "github.com/lib/pq"
)

func GetPostgresClient(config Configuration) *ent.Client {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=disable", config.Postgres.Broker, config.Postgres.Port, config.Postgres.User, config.Postgres.Password)
	fmt.Println(conn)
	client, err := ent.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func MigratePostgres(config Configuration) {
	client := GetPostgresClient(config)
	defer client.Close()

	err := client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func InsertNewCheckin(ctx context.Context, checkin *ent.CheckIn, config Configuration) (*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	ci, err := client.CheckIn.
		Create().
		SetAttendanceCode(checkin.AttendanceCode).
		SetStudentId(checkin.StudentId).
		SetStatus(checkin.Status).
		SetCheckinTime(checkin.CheckinTime).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("checkin was created: ", ci)
	return ci, nil
}

func GetCheckInByIdFromPostgres(ctx context.Context, id int64, config Configuration) (*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	ci, err := client.CheckIn.
		Query().
		Where(checkin.ID(int(id))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying checkin: %w", err)
	}
	log.Println("checkin returned: ", ci)
	return ci, nil
}

func GetCheckIns(ctx context.Context, config Configuration) ([]*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	cis, err := client.CheckIn.Query().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all checkins: %w", err)
	}

	return cis, nil
}

func GetByAttendanceCode(ctx context.Context, attendanceCode int64, config Configuration) ([]*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	cis, err := client.CheckIn.
		Query().
		Where(checkin.AttendanceCode(attendanceCode)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying checkins with attendence code: %w", err)
	}
	return cis, nil
}

func GetByStudentId(ctx context.Context, studentId string, config Configuration) ([]*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	cis, err := client.CheckIn.
		Query().Where(checkin.StudentId(studentId)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying checkins with studentId: %w", err)
	}
	return cis, nil
}

func GetByTime(ctx context.Context, from int64, to int64, config Configuration) ([]*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	defer client.Close()

	cis, err := client.CheckIn.
		Query().Where(checkin.CheckinTimeGTE(from), checkin.CheckinTimeLTE(to)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying checkins with studentId: %w", err)
	}
	return cis, nil
}
