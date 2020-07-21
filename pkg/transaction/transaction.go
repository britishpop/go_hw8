package transaction

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const formatDate = time.RFC1123

type Transaction struct {
	XMLName string    `xml:"transactions"`
	Id      int64     `xml:"id"`
	Type    string    `xml:"type"`
	Sum     int64     `xml:"sum"`
	Status  string    `xml:"status"`
	MCC     string    `xml:"mcc"`
	Date    time.Time `xml:"date"`
}

type Transactions struct {
	XMLName      string `xml:"transactions"`
	Transactions []*Transaction
}

func MakeTransactions(count int) []*Transaction {
	transactions := make([]*Transaction, count)
	for index := range transactions {
		v := &Transaction{
			`xml:"transaction"`,
			int64(index),
			"transfer",
			1000,
			"in progress",
			"4921",
			time.Date(2020, time.January, index, 11, 15, 10, 0, time.UTC),
		}
		transactions[index] = v
	}
	return transactions
}

func writeToFile(f io.Writer, transactions []*Transaction) error {
	mu := sync.Mutex{}
	records := [][]string{}
	for _, t := range transactions {
		record := []string{
			strconv.FormatInt(t.Id, 10),
			t.Type,
			strconv.FormatInt(t.Sum, 10),
			t.Status,
			t.MCC,
			t.Date.Format(formatDate),
		}
		mu.Lock()
		records = append(records, record)
		mu.Unlock()
	}

	w := csv.NewWriter(f)
	return w.WriteAll(records)
}

func ImportCSV(filename string) error {
	mu := sync.Mutex{}
	transactions := make([]*Transaction, 0)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return err
	}

	reader := csv.NewReader(bytes.NewReader(data))
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, row := range rows {

		mu.Lock()
		t, err := MapRowToTransaction(row)
		if err != nil {
			log.Println(err)
			return err
		}
		transactions = append(transactions, t)
		mu.Unlock()
	}
	return nil
}

func MapRowToTransaction(row []string) (*Transaction, error) {
	id, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sum, err := strconv.ParseInt(row[2], 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	date, err := time.Parse(formatDate, row[5])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tr := &Transaction{
		Id:     id,
		Type:   row[1],
		Sum:    sum,
		Status: row[3],
		MCC:    row[4],
		Date:   date,
	}
	return tr, nil
}

func ExportCSV(filename string, transactions []*Transaction) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}(file)

	err = writeToFile(file, transactions)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (t *Transactions) ExportXML(filename string) error {
	encodedXML, err := xml.Marshal(t)
	if err != nil {
		log.Print(err)
		return err
	}
	encodedXML = append([]byte(xml.Header), encodedXML...)

	file, err := os.Create(filename)
	if err != nil {
		log.Print(err)
		return err
	}
	defer file.Close()
	file.Write(encodedXML)

	return nil
}

func (t *Transactions) ImportXML(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return err
	}

	err = xml.Unmarshal(data, &t.Transactions)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
