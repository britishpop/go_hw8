package main

import (
	"go_hw_8/pkg/transaction"
	"log"
)

func main() {
	transactions := transaction.MakeTransactions(10)
	transaction.ExportCSV("example_import.csv", transactions)

	tr, errE := transaction.ImportCSV("example_import.csv")

	if errE == nil {
		for k, v := range tr {
			log.Printf("%d Transaction is %v \r\n", k, v)
		}
	}
}
