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
	Name       string `uri:"name"`
	ID         uint   `uri:"id"`
	Categories []Category
}

type Category struct {
	Title    string
	BudgetID uint
	ID       uint
	Total    float64
}

type Allocation struct {
	Amount        float64
	CategoryID    uint
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
	Name         string
	ID           uint
	BudgetID     uint
	Transactions []Transaction
}

func toExternalBudget(budget persistance.Budget, internalCategories []persistance.Category) Budget {
	var categories []Category
	for _, intCat := range internalCategories {
		categories = append(categories, toExternalCategory(intCat))
	}
	return Budget{
		Name:       budget.Name,
		ID:         budget.ID,
		Categories: categories,
	}
}

func toInternalBudget(budget Budget) persistance.Budget {
	return persistance.Budget{
		Name:  budget.Name,
		Model: gorm.Model{ID: budget.ID},
	}
}

func toExternalAccount(account persistance.Account, internalTransactions []persistance.Transaction) Account {
	var transactions []Transaction
	for _, intTransaction := range internalTransactions {
		transactions = append(transactions, toExternalTransaction(intTransaction))
	}
	return Account{
		ID:           account.ID,
		Name:         account.Name,
		BudgetID:     account.BudgetID,
		Transactions: transactions,
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

func toExternalAllocation(transaction persistance.Allocation) Allocation {
	return Allocation{
		Amount:        transaction.Amount,
		CategoryID:    transaction.CategoryID,
		ID:            transaction.ID,
		AssignedMonth: transaction.AssignedMonth,
	}
}

func toInternalAllocation(transaction Allocation) persistance.Allocation {
	return persistance.Allocation{
		Model:         gorm.Model{ID: transaction.ID},
		Amount:        transaction.Amount,
		CategoryID:    transaction.CategoryID,
		AssignedMonth: transaction.AssignedMonth,
	}
}
