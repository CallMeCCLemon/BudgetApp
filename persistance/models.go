package persistance

import (
	"net/http"
)

type Budget struct {
	Name       string
	Categories []string
	Accounts   []string
	Id         string
}

type Category struct {
	Title          string
	AllocatedFunds float64
	BudgetID       string
	ID             string
	Total          float64
	Allocations    []string
}

type Allocation struct {
	// date
	Amount     float64
	CategoryID string
	ID         string
}

type Transaction struct {
	Amount   float64
	Memo     string
	Account  Account
	Category Category
	ID       string
	// date
}

type Account struct {
	Name string
	ID   string
}

type StorageDao struct {
	Connection http.Client
}

func (*StorageDao) ReadBudget(id string) (budget Budget, err error) {
	budget = Budget{}
	// Get IDs from Storage layer.

	return budget, nil
}

func (*StorageDao) WriteBudget(budget Budget) (id string, err error) {
	return "", nil
}

func (*StorageDao) ReadCategory(id string) (category Category, err error) {
	return Category{}, nil
}

func (*StorageDao) WriteCategory(category Category) (id string, err error) {
	return "", nil
}

func (*StorageDao) ReadAccount() (account Account, err error) {
	return Account{}, nil
}

func (*StorageDao) WriteAccount(account Account) (id string, err error) {
	return "", nil
}

func (*StorageDao) ReadTransaction() (transaction Transaction, err error) {
	return Transaction{}, nil
}

func (*StorageDao) WriteTransaction(transaction Transaction) (id string, err error) {
	return "", nil
}

func (*StorageDao) GetAllocation(date string) (allocation Allocation, err error) {
	return Allocation{}, nil
}
