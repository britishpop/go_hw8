package main

import (
	"fmt"
	"go_hw_8/pkg/transaction"
	"os"
	"time"
)

func main() {
	transactions := transaction.MakeTransactions(10)
	if err := transaction.ExportCSV("example_import.csv", transactions); err != nil {
		os.Exit(10)
	}

	if errE := transaction.ImportCSV("example_import.csv"); errE != nil {
		os.Exit(10)
	}

	d := time.Date(2020, time.January, 2, 11, 15, 10, 0, time.UTC)
	a := d.Format(time.RFC3339)
	fmt.Println(a)
}
