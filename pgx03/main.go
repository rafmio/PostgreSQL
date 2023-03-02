package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	log.Println("starting program")

	databaseUrl := "postgres://postgres:qwq121@localhost:5432/budgeting"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//to close DB pool
	defer dbPool.Close()
}

func AddTransaction(dpPool *pgxpool.Pool) {
	log.Println("adding new entry")
	entry := &Transaction{}

	fmt.Printf("Account: ")
	fmt.Scanf("%s", entry.account)

	fmt.Printf("Date of transaction: ")
	fmt.Scanf("%s", entry.trdate)

	fmt.Printf("Type of document: ")
	fmt.Scanf("%s", entry.trtype)

	fmt.Printf("Date of document: ")
	fmt.Scanf("%s", entry.docdate)

	fmt.Printf("Number of document: ")
	fmt.Scanf("%s", entry.docnum)

	fmt.Printf("Counterparty: ")
	fmt.Scanf("%s", entry.counterparty)

	fmt.Printf("Tax ID of counterparty: ")
	fmt.Scanf("%s", entry.cntp_tax_id)

	fmt.Printf("Contract: ")
	fmt.Scanf("%s", entry.cntp_contract)

	fmt.Printf("Purpose of payment: ")
	fmt.Scanf("%s", entry.purpose)

	fmt.Printf("Direction (inflow/outflow): ")
	fmt.Scanf("%s", entry.direction)

	fmt.Printf("Amount: ")
	fmt.Scanf("%.2f", entry.amount)

	fmt.Printf("Item: ")
	fmt.Scanf("%s", entry.item)

	fmt.Printf("Comment: ")
	fmt.Scanf("%s", entry.comment)

	queryText := fmt.Sprintf("INSERT INTO transactions(account, trdate, trtype, docdate, docnumb, counterpary, cntp_tax_id, cntp_contract, purpose, comment, direction, amount, item) VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %.2f, %s)",
		entry.account, entry.trdate, entry.trtype, entry.docdate, entry.docnumb, entry.counterparty, entry.cntp_tax_id, entry.cntp_contract, entry.purpose, entry.comment, entry.direction, entry.item)

}
