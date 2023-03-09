package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// postgres://postgres:qwq121@localhost:5432/ethcontract

type AccountEntry struct {
	AccNumber string
	Password  string
}

func InsertAccount(entry AccountEntry) {
	config, err := pgx.ParseConfig("postgresql://postgres:qwq121@localhost:5432/ethcontract?sslmode=disable")
	if err != nil {
		fmt.Println("parsing config: ", err.Error())
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("connecting to database: ", err.Error())
	}
	defer conn.Close(context.Background())

	acc := entry.AccNumber
	psw := entry.Password

	query := fmt.Sprintf("INSERT INTO accounts (accnumber, password) VALUES('%s', '%s');", acc, psw)

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("adding entry: ", err.Error())
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	}

	fmt.Println("Database connection closed")
}
