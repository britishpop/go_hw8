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

	transactionsXML := &transaction.Transactions{
		Transactions: transactions,
	}

	if err := transactionsXML.ExportXML("example_import_xml.xml"); err != nil {
		os.Exit(10)
	}

	transactionsXMLToImport := &transaction.Transactions{}
	if err := transactionsXMLToImport.ImportXML("example_import_xml.xml"); err != nil {
		os.Exit(10)
	}
}
