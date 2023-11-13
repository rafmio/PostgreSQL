package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func main() {
	// open up our database connection
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5432/test")

	// defer the close till after the main function has finished executing
	defer conn.Close(context.Background())

	var greeting string

	conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
}
