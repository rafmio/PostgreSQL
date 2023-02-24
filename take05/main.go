// https://dev.to/kushagra_mehta/postgresql-with-go-in-2021-3dfg
package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:qwq121@localhost:5432/test_db")
	if err != nil {
		fmt.Println("connecting to test_db: ", err.Error())
	} else {
		fmt.Println("connections to database test_db was esteblished")
	}
	defer conn.Close(context.Background())

	createTableQuery := "CREATE TABLE IF NOT EXISTS tablego1 (user_id serial PRIMARY KEY, username VARCHAR (50), email VARCHAR(255));"

	_, err = conn.Exec(context.Background(), createTableQuery)
	if err != nil {
		fmt.Println("creating table: ", err.Error())
	} else {
		fmt.Println("the sql table via Go was created")
	}
}
