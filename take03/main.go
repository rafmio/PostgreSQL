package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "qwq121"
	dbname = "test_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	insertStmt := `insert into "students"("name", "roll") values('Valera', 2)`
	_, e := db.Exec(insertStmt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

    // psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s 
	// sslmode=disable", host, port, user, password, dbname)
