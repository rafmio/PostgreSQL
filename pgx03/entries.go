package main

import (
	"reflect"
	"log"
)

// Структура транзакции
type Transaction struct {
	Account       string
	TrDate        string
	TrType        string
	DocDate       string
	DocNumb       string
	Counterparty  string
	CntpTaxID   string
	Contract  string
	Purpose       string
	Direction     string
	Amount        float64
	Item          string
	Comment       string
}

// Описание каждого поля структуры транзакции для вывода
var TransactionFieldsDescription map[string]string = map[string]string {
	"Account" : "Account",
	"TrDate": "Date of transaction",
	"TrType": "Type of transaction",
	"DocDate": "Date of document",
	"DocNumb": "Number of document",
	"Counterparty": "Counterparty",
	"CntpTaxID": "Tax ID of counterparty",
	"Countract": "Countract",
	"Purpose": "Purpose of payment",
	"Direction": "Direction of payment (inflow/outflow)",
	"Amount": "Amount of payment (roubles)",
	"Item": "Item of cost",
	"Comment": "Comment",
}

// Атомарная функция добавления транзации со всеми полями (не по отдельным полям)
func (transaction *Transaction) AddFullTransaction() err error {
	log.Println("adding full transaction")

	values := reflect.ValueOf(transaction)
	types := values.Type()

// Цикл заполнения структуры
	for i := 0; i < values.NumField(); i++ {
		fmt.Printf(TransactionFieldsDescription[types.Field(i).Name], ": ")

		// Значение перебираемого поля помещаем для удобства в отдельную
		// переменную, чтоб не вызывать лишний раз функции
		fieldValue := types.Field(i).Name

		// Все поля - string, кроме поля суммы (Amount float64)
		// Для поля Amount float64 отдельное сканирование с
		if fieldValue != "Amount" {
			fmt.Scanf("%s", transaction.fieldValue)
		} else {
			fmt.Scanf("%.2f", transaction.fieldValue)
		}
	}
}
