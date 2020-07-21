package main

import (
	"go_hw_8/pkg/transaction"
	"os"
)

func main() {
	transactions := transaction.MakeTransactions(10)
	if err := transaction.ExportCSV("example_import.csv", transactions); err != nil {
		os.Exit(10)
	}

	if errE := transaction.ImportCSV("example_import.csv"); errE != nil {
		os.Exit(10)
	}
}
