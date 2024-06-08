package setup

import (
	"BudgetingApp/persistance"
	"log"
	"os"
	"testing"
)

func Test_CreateTables(t *testing.T) {
	dao, err := persistance.NewStorageDao(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		"budgetapp")
	if err != nil {
		log.Fatal("Failed to connect to the DB! ", err)
	}
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Budget{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Account{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Allocation{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Transaction{})
	_ = dao.GormDB.Migrator().CreateTable(&persistance.Category{})

}
