package main

import (
	"go_hw_8/pkg/transaction"
	"log"
	"os"
)

func main() {
	tr := transaction.MakeTransactions(3)

	if err := transaction.ExportCSV("example_import.csv", tr); err != nil {
		os.Exit(10)
	}

	transactionsCSV, err := transaction.ImportCSV("example_import.csv")
	if err != nil {
		os.Exit(10)
	}
	for _, v := range transactionsCSV {
		log.Printf("New transaction from CSV file: %v", *v)
	}

	if err := transaction.ExportJSON("example_import_json.csv", tr); err != nil {
		os.Exit(10)
	}

	transactionsJSON, err := transaction.ImportJSON("example_import_json.csv")
	if err != nil {
		os.Exit(10)
	}
	for _, v := range transactionsJSON {
		log.Printf("New transaction from JSON file: %v", *v)
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
