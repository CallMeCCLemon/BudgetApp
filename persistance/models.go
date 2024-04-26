package persistance

import (
	"crypto/tls"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Budget struct {
	Name       string
	Categories []uuid.UUID
	Accounts   []uuid.UUID
	ID         uuid.UUID
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
	Date     time.Time
}

type Account struct {
	Name string
	ID   uuid.UUID
}

type StorageDao struct {
	DB *sql.DB
}

func NewStorageDao(username string, password string, address string, dbname string) (*StorageDao, error) {
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   address,
		DBName: dbname,
		//TLSConfig: "skip-verify",
		TLS: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS12,
		},
		ParseTime: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
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

func (dao *StorageDao) ReadBudgets() (budgets []Budget, err error) {
	log.Println("DB status: ", dao.DB.Stats())
	result, err := dao.DB.Query("SELECT * FROM Budgets")
	if err != nil {
		return []Budget{}, err
	}
	var joinedCategories string
	var joinedAccounts string

	for result.NextResultSet() {
		budgetReader := Budget{}
		err = result.Scan(&budgetReader.ID, &budgetReader.Name, &joinedCategories, &joinedAccounts)
		if err != nil {
			return nil, err
		}
		budgets = append(budgets, budgetReader)
	}

	return
}

func (dao *StorageDao) ReadBudget(id uuid.UUID) (budget *Budget, err error) {
	row := dao.DB.QueryRow("SELECT * FROM Budgets WHERE ID=?", id)
	var joinedCategories string
	var joinedAccounts string
	tmpBudget := Budget{}
	err = row.Scan(&tmpBudget.ID, &tmpBudget.Name, &joinedCategories, &joinedAccounts)
	if err != nil {
		return nil, err
	}
	accounts, err := toUUIDs(strings.Split(joinedAccounts, ","))
	if err != nil {
		return nil, err
	}
	tmpBudget.Accounts = accounts

	categories, err := toUUIDs(strings.Split(joinedCategories, ","))
	if err != nil {
		return nil, err
	}
	tmpBudget.Categories = categories
	budget = &tmpBudget
	return
}

func (dao *StorageDao) WriteBudget(budget Budget) (id *uuid.UUID, err error) {
	result, err := dao.DB.Exec("INSERT INTO Budgets (Name, Categories, Accounts, Id) VALUES (?, ?, ?, ?)",
		budget.Name, strings.Join(toStrings(budget.Categories), ","), strings.Join(toStrings(budget.Accounts), ","), budget.ID)
	if err != nil {
		return nil, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id = &budget.ID
	return
}

func (dao *StorageDao) DeleteBudget(id uuid.UUID) (deletedID *uuid.UUID, err error) {
	_, err = dao.DB.Exec("DELETE FROM Budgets WHERE ID=?", id)
	if err != nil {
		return nil, err
	}
	deletedID = &id
	return
}

func (dao *StorageDao) ReadCategory(id uuid.UUID) (category Category, err error) {
	row := dao.DB.QueryRow("SELECT * FROM Categories WHERE ID=?", id)
	var allocations string
	err = row.Scan(&category.ID, &category.Title, &category.AllocatedFunds, &category.Total, &allocations)
	category.Allocations = strings.Split(allocations, ",")
	if err != nil {
		return Category{}, err
	}
	return
}

func (dao *StorageDao) WriteCategory(category Category) (id *uuid.UUID, err error) {
	result, err := dao.DB.Exec("INSERT INTO Categories (ID, Title, AllocatedFunds, Total, Allocations) VALUES (?, ?, ?, ?, ?)",
		category.ID, category.Title, category.AllocatedFunds, category.Total, strings.Join(category.Allocations, ","))
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

func (dao *StorageDao) DeleteCategory(id uuid.UUID) (deletedID *uuid.UUID, err error) {
	_, err = dao.DB.Exec("DELETE FROM Categories WHERE ID=?", id)
	if err != nil {
		return nil, err
	}

	deletedID = &id
	return
}

func (dao *StorageDao) ReadAccount(id uuid.UUID) (account *Account, err error) {
	row := dao.DB.QueryRow("SELECT * FROM Accounts WHERE ID=?", id)
	var tempAccount Account

	err = row.Scan(&tempAccount.ID, &tempAccount.Name)

	if err != nil {
		return nil, err
	}

	account = &tempAccount
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

func (dao *StorageDao) DeleteAccount(id uuid.UUID) (deletedID *uuid.UUID, err error) {
	_, err = dao.DB.Exec("DELETE FROM Accounts WHERE ID=?", id)
	if err != nil {
		return nil, err
	}

	deletedID = &id
	return
}

func (dao *StorageDao) ReadTransaction(id uuid.UUID) (transaction *Transaction, err error) {
	log.Default().Printf("UUID: %s", id.String())
	row := dao.DB.QueryRow("SELECT * FROM Transactions WHERE ID=?", id)

	var categoryID uuid.UUID
	var accountID uuid.UUID

	var tempTransaction Transaction

	err = row.Scan(&tempTransaction.ID, &tempTransaction.Amount, &tempTransaction.Memo, &accountID, &categoryID, &tempTransaction.Date)
	if err != nil {
		return nil, err
	}
	account, err := dao.ReadAccount(accountID)
	tempTransaction.Account = *account
	if err != nil {
		return nil, err
	}

	tempTransaction.Category, err = dao.ReadCategory(categoryID)
	if err != nil {
		return nil, err
	}
	transaction = &tempTransaction
	return
}

func (dao *StorageDao) WriteTransaction(transaction Transaction) (id uuid.UUID, err error) {
	log.Default().Printf("UUID: %s", transaction.ID.String())
	result, err := dao.DB.Exec("INSERT INTO Transactions (amount, memo, accountID, categoryID, id, Date) VALUES (?, ?, ?, ?, ?, ?)", transaction.Amount, transaction.Memo, transaction.Account.ID, transaction.Category.ID, transaction.ID, transaction.Date)
	if err != nil {
		return uuid.UUID{}, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return uuid.UUID{}, err
	}
	id = transaction.ID
	return
}

func (dao *StorageDao) DeleteTransaction(id uuid.UUID) (deletedID *uuid.UUID, err error) {
	_, err = dao.DB.Exec("DELETE FROM Transactions WHERE ID=?", id)
	if err != nil {
		return nil, err
	}

	deletedID = &id
	return
}

func (dao *StorageDao) GetAllocation(date string) (allocation Allocation, err error) {
	return Allocation{}, nil
}

func toStrings(ids []uuid.UUID) []string {
	var stringIds []string
	for _, id := range ids {
		stringIds = append(stringIds, id.String())
	}
	return stringIds
}

func toUUIDs(ids []string) ([]uuid.UUID, error) {
	var uuids []uuid.UUID
	for _, id := range ids {
		parsedID, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, parsedID)
	}
	return uuids, nil
}
