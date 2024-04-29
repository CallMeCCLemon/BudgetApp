package service

import (
	"BudgetingApp/persistance"
	"gorm.io/gorm"
	"time"
)

type Query struct {
	ID uint `uri:"id"`
}

type Budget struct {
	Name string `uri:"name"`
	ID   uint   `uri:"id"`
}

type Category struct {
	Title    string
	BudgetID uint
	ID       uint
	Total    float64
}

type Allocation struct {
	// date
	Amount        float64
	CategoryID    string
	ID            uint
	AssignedMonth time.Time
}

type Transaction struct {
	Amount     float64
	Memo       string
	AccountID  uint
	CategoryID uint
	ID         uint
	Date       time.Time
}

type Account struct {
	Name     string
	ID       uint
	BudgetID uint
}

func toExternalBudget(budget persistance.Budget) Budget {
	return Budget{
		Name: budget.Name,
		ID:   budget.ID,
	}
}

func toInternalBudget(budget Budget) persistance.Budget {
	return persistance.Budget{
		Name:  budget.Name,
		Model: gorm.Model{ID: budget.ID},
	}
}

func toExternalAccount(account persistance.Account) Account {
	return Account{
		ID:       account.ID,
		Name:     account.Name,
		BudgetID: account.BudgetID,
	}
}

func toInternalAccount(account Account) persistance.Account {
	return persistance.Account{
		Model:    gorm.Model{ID: account.ID},
		Name:     account.Name,
		BudgetID: account.BudgetID,
	}
}

func toExternalCategory(category persistance.Category) Category {
	return Category{
		Title:    category.Title,
		BudgetID: category.BudgetID,
		ID:       category.ID,
		Total:    category.Total,
	}
}

func toInternalCategory(category Category) persistance.Category {
	return persistance.Category{
		Model:    gorm.Model{ID: category.ID},
		Title:    category.Title,
		BudgetID: category.BudgetID,
		Total:    category.Total,
	}
}

func toExternalTransaction(transaction persistance.Transaction) Transaction {
	return Transaction{
		Amount:     transaction.Amount,
		Memo:       transaction.Memo,
		AccountID:  transaction.AccountID,
		CategoryID: transaction.CategoryID,
		ID:         transaction.ID,
		Date:       transaction.Date,
	}
}

func toInternalTransaction(transaction Transaction) persistance.Transaction {
	return persistance.Transaction{
		Model:      gorm.Model{ID: transaction.ID},
		Amount:     transaction.Amount,
		Memo:       transaction.Memo,
		AccountID:  transaction.AccountID,
		CategoryID: transaction.CategoryID,
		Date:       transaction.Date,
	}
}
