package main

import (
	"BudgetingApp/persistance"
	"github.com/google/uuid"
	"log"
	"os"
)

func main() {

	// Routes
	// /budget -> Get all budgets for user
	// /budget/${id} -> Get single budget with categories and computations
	// /account/${id} -> Get account and transactions
	//
	//

	dao, err := persistance.NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	_, err = dao.WriteAccount(persistance.Account{
		Name: "test-account-001",
		ID:   uuid.New(),
	})
	if err != nil {
		log.Fatal("Failed to write account to the MySQL DB!", err)
	}

}

type Budget struct {
	Name       string
	Categories []Category
	Accounts   []Account
	ID         uuid.UUID
}

type Category struct {
	Title       string
	Budget      Budget
	ID          uuid.UUID
	Total       float64
	Allocations []Allocation
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
	// date
}

type Account struct {
	Name string
	ID   uuid.UUID
}

func getBudget(storageDao *persistance.StorageDao, id string, month string) (budget Budget, err error) {
	internalBudget, err := storageDao.ReadBudget(id)
	if err != nil {
		return Budget{}, err
	}
	var categories []Category
	var errors []error
	for _, categoryId := range internalBudget.Categories {
		category, err := storageDao.ReadCategory(categoryId)
		if err != nil {
			errors = append(errors, err)
		}
		categories = append(categories, Category{
			Title: category.Title,
			//Allocations:
		})
	}
	budget = Budget{
		Name: internalBudget.Name,
		//Categories: internalBudget.Categories,

	}

	return
}
