package main

// Структура транзакции
type Transaction struct {
	Account      string
	TrDate       string
	TrType       string
	DocDate      string
	DocNumb      string
	Counterparty string
	CntpTaxID    string
	Contract     string
	Purpose      string
	Direction    string
	Amount       float64
	Item         string
	Comment      string
}

// Описание каждого поля структуры транзакции для вывода
var TransactionFieldsDescription map[string]string = map[string]string{
	"Account":      "Account",
	"TrDate":       "Date of transaction",
	"TrType":       "Type of transaction",
	"DocDate":      "Date of document",
	"DocNumb":      "Number of document",
	"Counterparty": "Counterparty",
	"CntpTaxID":    "Tax ID of counterparty",
	"Countract":    "Countract",
	"Purpose":      "Purpose of payment",
	"Direction":    "Direction of payment (inflow/outflow)",
	"Amount":       "Amount of payment (roubles)",
	"Item":         "Item of cost",
	"Comment":      "Comment",
}

type SetTransaction interface {
	SetAccount()
	SetTrDate()
	SetTrType()
	SetDocDate()
	SetDocNumb()
	SetCounterparty()
	SetCntpTaxID()
	SetContract()
	SetPurpose()
	SetDirection()
	SetAmount()
	SetItem()
	SetComment()
}

func (transaction *Transaction) SetAccount() {
	log.Println("setting ", TransactionFieldsDescription["Account"])
	fmt.Printf(TransactionFieldsDescription["Account"], ": ")
	fmt.Scanf("%s", transaction.Account)
}

func (transaction *Transaction) SetTrDate() {
	log.Println("settig ", TransactionFieldsDescription["TrDate"])
	fmt.Printf(TransactionFieldsDescription["TrDate"], ": ")
	fmt.Scanf("%s", transaction.TrDate)
}

func 