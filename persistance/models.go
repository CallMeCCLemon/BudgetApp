package persistance

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Budget struct {
	gorm.Model
	Name string
}

type Category struct {
	gorm.Model
	Title    string
	BudgetID uint `gorm:"foreignKey:BudgetRefer"`
	Total    float64
}

type Allocation struct {
	gorm.Model
	Amount        float64
	CategoryID    uint
	AssignedMonth time.Time
}

type Transaction struct {
	gorm.Model
	Amount     float64
	Memo       string
	AccountID  uint `gorm:"foreignKey:AccountRefer"`
	CategoryID uint `gorm:"foreignKey:CategoryRefer"`
	Date       time.Time
}

type Account struct {
	gorm.Model
	Name     string
	BudgetID uint `gorm:"foreignKey:BudgetRefer"`
}

type StorageDao struct {
	GormDB *gorm.DB
}

func NewStorageDao(username string, password string, address string, port string, dbname string) (*StorageDao, error) {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", address, username, password, dbname, port),
	}

	psql := postgres.New(cfg)
	gormCfg := gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	gormDB, err := gorm.Open(psql, &gormCfg)
	if err != nil {
		return nil, err
	}

	return &StorageDao{GormDB: gormDB}, nil
}

func (dao *StorageDao) GetAllBudgets() (budgets []Budget, err error) {
	dao.GormDB.Find(&budgets)
	return
}

func (dao *StorageDao) GetAllTransactions() (transactions []Transaction, err error) {
	dao.GormDB.Find(&transactions)
	return
}

func (dao *StorageDao) GetBudget(id uint) (budget *Budget, err error) {
	budget = &Budget{
		Model: gorm.Model{ID: id},
	}
	result := dao.GormDB.First(&budget)
	if result.Error != nil {
		return budget, result.Error
	}
	if result.RowsAffected == 0 {
		return budget, errors.New(fmt.Sprintf("No rows found for ID %d", id))
	}
	return
}

func (dao *StorageDao) GetAccount(id uint) (account *Account, err error) {
	account = &Account{
		Model: gorm.Model{ID: id},
	}
	result := dao.GormDB.First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("No rows found for ID %d", id))
	}
	return
}

func (dao *StorageDao) GetTransactionsForAccount(accountID uint) (transactions []Transaction, err error) {
	result := dao.GormDB.Where("account_id = ?", accountID).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}

func (dao *StorageDao) GetCategoriesForBudget(budgetId uint) (categories []Category, err error) {
	result := dao.GormDB.Where("budget_id = ?", budgetId).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}

func (dao *StorageDao) GetCategory(id uint) (category *Category, err error) {
	category = &Category{
		Model: gorm.Model{ID: id},
	}
	result := dao.GormDB.First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("No rows found for ID %d", id))
	}
	return
}

func (dao *StorageDao) GetTransaction(id uint) (transaction *Transaction, err error) {
	transaction = &Transaction{
		Model: gorm.Model{ID: id},
	}
	result := dao.GormDB.First(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("No rows found for ID %d", id))
	}
	return
}

func (dao *StorageDao) GetAllocation(id uint) (allocation *Allocation, err error) {
	allocation = &Allocation{
		Model: gorm.Model{ID: id},
	}
	result := dao.GormDB.First(&allocation)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("No rows found for ID %d", id))
	}
	return
}
