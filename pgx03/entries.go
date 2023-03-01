package main

type Transaction struct {
	account       string
	trdate        string
	trtype        string
	docdate       string
	docnumb       string
	counterparty  string
	cntp_tax_id   string
	cntp_contract string
	purpose       string
	comment       string
	direction     string
	amount        int
	item          string
}
