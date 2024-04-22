package persistance

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Budget struct {
	Name       string
	Categories []string
	Accounts   []string
	Id         uuid.UUID
}

type Category struct {
	Title          string
	AllocatedFunds float64
	BudgetID       uuid.UUID
	ID             uuid.UUID
	Total          float64
	Allocations    []string
}

type Allocation struct {
	// date
	Amount     float64
	CategoryID string
	ID         uuid.UUID
}

type Transaction struct {
	Amount   float64
	Memo     string
	Account  Account
	Category Category
	ID       uuid.UUID
	date     time.Time
}

type Account struct {
	Name string
	ID   uuid.UUID
}

type StorageDao struct {
	DB *sql.DB
}

func NewStorageDao(username string, password string, address string, dbname string) (*StorageDao, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s`, username, password, address, dbname))
	if err != nil {
		return nil, err
	}
	//// TODO: Figure out if this is how this is intended to be done.
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(db)
	return &StorageDao{DB: db}, nil
}

func (dao *StorageDao) ReadBudget(id string) (budget Budget, err error) {
	budget = Budget{}

	// Get IDs from Storage layer.

	return budget, nil
}

func (dao *StorageDao) WriteBudget(budget Budget) (id string, err error) {
	return "", nil
}

func (dao *StorageDao) ReadCategory(id string) (category Category, err error) {
	return Category{}, nil
}

func (dao *StorageDao) WriteCategory(category Category) (id *uuid.UUID, err error) {
	result, err := dao.DB.Exec("INSERT INTO Categories (ID, Title, AllocatedFunds, BudgetID, Total, Allocations) VALUES (?, ?, ?, ?, ?, ?)",
		category.ID, category.Title, category.AllocatedFunds, category.BudgetID, category.Total, category.Allocations)
	if err != nil {
		return nil, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id = &category.ID
	return
}

func (dao *StorageDao) ReadAccount(id uuid.UUID) (account Account, err error) {
	row := dao.DB.QueryRow("SELECT * FROM Accounts WHERE ID=?", id)
	if err != nil {
		return Account{}, err
	}

	err = row.Scan(&account.ID, &account.Name)
	if err != nil {
		return Account{}, err
	}
	return
}

func (dao *StorageDao) WriteAccount(account Account) (id *uuid.UUID, err error) {
	result, err := dao.DB.Exec("INSERT INTO Accounts (ID, Name) VALUES (?, ?)", account.ID, account.Name)
	if err != nil {
		return nil, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id = &account.ID
	return
}

func (dao *StorageDao) ReadTransaction() (transaction Transaction, err error) {
	return Transaction{}, nil
}

func (dao *StorageDao) WriteTransaction(transaction Transaction) (id *uuid.UUID, err error) {
	result, err := dao.DB.Exec("INSERT INTO Transactions (amount, memo, accountID, categoryID, id) VALUES (?, ?, ?, ?, ?)", transaction.Amount, transaction.Memo, transaction.Account, transaction.Category, transaction.ID)
	if err != nil {
		return nil, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id = &transaction.ID
	return
}

func (dao *StorageDao) GetAllocation(date string) (allocation Allocation, err error) {
	return Allocation{}, nil
}
