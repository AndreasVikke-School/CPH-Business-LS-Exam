package main

import (
	"context"
	"fmt"
	"log"

	"postgres_service/ent"
	"postgres_service/ent/checkin"

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
	err := client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	CreateUser(context.Background(), client)
	defer client.Close()
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.CheckIn, error) {
	ci, err := client.CheckIn.
		Create().
		SetAttendanceCode(1234567).
		SetStudentId(1).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("checkin was created: ", ci)
	return ci, nil
}

func GetCheckInByIdFromPostgres(ctx context.Context, id int64, config Configuration) (*ent.CheckIn, error) {
	client := GetPostgresClient(config)
	ci, err := client.CheckIn.
		Query().
		Where(checkin.ID(int(id))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("checkin returned: ", ci)
	defer client.Close()
	return ci, nil
}
