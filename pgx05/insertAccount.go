// https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountEntry struct {
	AccNumber string // 40 symbols
	Password  string // 256 symbols
}

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

var connString string = "postgres://postgres:qwq121@localhost:5432/ethcontract"

func NewPG(ctx context.Context, connString string) (*postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			return fmt.Errorf("unable to establish connection pool: %w", err)
		}
		pgInstance = &postgres{db}
	})

	return pgInstance, nil
}

func ConnectDB(entry AccountEntry) {
	fmt.Println("connection to database")

	// establish a connection with specified DB
	dbPool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	// execute insert query
	ExecuteInsertAccountQuery(dbPool, entry)
}

func ExecuteInsertAccountQuery(dbPool *pgxpool.Pool, entry AccountEntry) {
	fmt.Println("starting executing of insert query")
	acc := entry.AccNumber
	psw := entry.Password
	SQLQuery := fmt.Sprintf("INSERT INTO accounts (accnumber, password) VALUES('%s', '%s');", acc, psw)
	rows, err := dbPool.Query(context.Background(), SQLQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error inserting data to DB", err)
	}

	fmt.Println("wrote: ", rows)
}
