package main

func main() {
	accountNumber := "352457987298754320"
	AESPassword := "a87d7f9eue9f9d7d87fd8783f439rf398f8934r82f948fg"
	var newAccount = AccountEntry{AccNumber: accountNumber, Password: AESPassword}

	ConnectDB(newAccount)
}
