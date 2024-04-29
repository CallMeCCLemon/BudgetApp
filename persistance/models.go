package persistance

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"

	msql "gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"
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
	AssignedMonth datatypes.Date
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
	DB     *sql.DB
	GormDB *gorm.DB
}

func NewStorageDao(username string, password string, address string, dbname string) (*StorageDao, error) {
	loc, err := time.LoadLocation("UTC")
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
		Loc:       loc,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	gormDb, err := createGormInstance(db)
	//// TODO: Figure out if this is how this is intended to be done.
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(db)
	return &StorageDao{DB: db, GormDB: gormDb}, nil
}

func createGormInstance(db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(msql.New(msql.Config{
		Conn: db,
	}), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}

func (dao *StorageDao) GetAllBudgets() (budgets []Budget, err error) {
	dao.GormDB.Find(&budgets)
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
