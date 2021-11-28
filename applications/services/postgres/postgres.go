package main

import (
	"context"
	"fmt"
	"log"

	"redis_service/ent"

	_ "github.com/lib/pq"
)

func MigratePostgres(config Configuration) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s", config.Postgres.Broker, config.Postgres.Port, config.Postgres.User, config.Postgres.Password)
	fmt.Println(conn)
	client, err := ent.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Run the auto migration tool.
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// func GetCheckInFromPostgres(code int64, config Configuration) (int64, int64, error) {

// }
