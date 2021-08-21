package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func db_connect() *pgx.Conn {

	db_uri := os.Getenv("DATABASE_URI")
	conn, err := pgx.Connect(context.Background(), db_uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
