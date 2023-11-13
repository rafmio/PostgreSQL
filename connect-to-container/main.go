package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	// "github.com/jackc/pgx/v5"
)

func main() {
	config, err := pgx.ParseConfig("postgres://postgres:password@localhost:5432/logentriesdb")
	if err != nil {
		log.Fatal("Failed to parse connection string: ", err.Error())
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stdout, "conn.string was parsed\n")
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Unable to connect to database:", err.Error())
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stdout, "Connection has been established\n")
	}
	defer conn.Close(context.Background())

	var result int
	err = conn.QueryRow(context.Background(), "SELECT 1+1").Scan(&result)
	if err != nil {
		log.Fatal("Query failed", err)
	}
	fmt.Fprintf(os.Stdout, "Result: %s\n", result)
}

// ENV POSTGRES_USER postgres
// ENV POSTGRES_PASSWORD password
// ENV POSTGRES_DB logentriesdb

// CREATE TABLE logentries
