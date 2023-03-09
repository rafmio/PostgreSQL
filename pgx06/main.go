package main

import "fmt"

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "3002999848"
	AESPassword := "8h8vr57t7dk94d76r9td8783faa8t7598f89gh482f9487t6r"

	// Assign to struct variable values of account number and password
	var newAccount = AccountEntry{AccNumber: accountNumber, Password: AESPassword}

	// Connecting to DB
	InsertAccount(newAccount)

	fmt.Println("The entry account number with password was added to DB")
}
