package persistance

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// ParseCsvTransactions Expecting rows to contain FOUR fields in the order "Date,Amount,Memo,CategoryID"!
func ParseCsvTransactions(file *os.File, AccountID uint) (transactions []Transaction, err error) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close CSV File!")
		}
	}(file)

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV file", err)
		return []Transaction{}, err
	}

	for _, record := range records {
		newTransaction, err := csvRowToTransaction(record, AccountID)
		if err != nil {
			log.Fatal("Failed to parse record as transaction!", record)
			return []Transaction{}, err
		}
		transactions = append(transactions, *newTransaction)
	}

	return transactions, nil
}

// csvRowToTransaction Expecting row to contain FOUR fields in the order "Date,Amount,Memo,CategoryID"!
func csvRowToTransaction(row []string, AccountID uint) (*Transaction, error) {
	const layout = "01/02/2006"

	date, err := time.Parse(layout, row[0])
	if err != nil {
		log.Fatal("Failed to parse Date as MM/DD/YYYY Format!!", row[0])
		return nil, err
	}

	amount, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		log.Fatal("Failed to parse Amount as Float!", row[1])
		return nil, err
	}

	categoryID, err := strconv.ParseUint(row[3], 0, 0)
	if err != nil {
		log.Fatal("Failed to parse Category as uint!", row[3])
		return nil, err
	}

	return &Transaction{
		Amount:     amount,
		Memo:       row[2],
		AccountID:  AccountID,
		CategoryID: uint(categoryID),
		Date:       date,
	}, nil
}
