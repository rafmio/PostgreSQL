// https://dev.to/kushagra_mehta/postgresql-with-go-in-2021-3dfg
package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:qwq121@localhost:5432/test_db")
	defer conn.Close(context.Background())

	var greeting string

	conn.QueryRow(context.Background(), "select 'Hello-mello, world-mworld!'").Scan(&greeting)
	fmt.Println(greeting)
}
