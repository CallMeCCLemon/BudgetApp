package setup

import (
	"BudgetingApp/persistance"
	"log"
	"os"
	"testing"
)

func Test_CreateTables(t *testing.T) {
	dao, err := persistance.NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Budget{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Account{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Allocation{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Transaction{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Category{})
}
