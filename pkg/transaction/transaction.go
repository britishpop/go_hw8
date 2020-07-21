package transaction

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
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
	Id     int64
	Type   string
	Sum    int64
	Status string
	MCC    string
	Date   time.Time
}

func MakeTransactions(count int) []*Transaction {
	transactions := make([]*Transaction, count)
	for index := range transactions {
		v := &Transaction{
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

func ImportCSV(filename string) ([]*Transaction, error) {
	mu := sync.Mutex{}
	transactions := make([]*Transaction, 0)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(data))
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, row := range rows {

		mu.Lock()
		t, err := MapRowToTransaction(row)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		transactions = append(transactions, t)
		mu.Unlock()
	}
	return transactions, nil
}

func ExportJSON(filename string, transactions []*Transaction) error {
	encodedJson, err := json.Marshal(transactions)
	if err != nil {
		log.Println(err)
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	file.Write(encodedJson)

	return nil
}

func ImportJSON(filename string) ([]*Transaction, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	transactions := make([]*Transaction, 0)

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return transactions, nil
}
