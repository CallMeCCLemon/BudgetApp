package persistance

import (
	"log"
	"os"
	"testing"

	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStorageDao_Budgets(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"
	categories := []uuid.UUID{uuid.New(), uuid.New()}
	accounts := []uuid.UUID{uuid.New(), uuid.New()}

	budget := Budget{
		Name:       name,
		ID:         id,
		Categories: categories,
		Accounts:   accounts,
	}

	t.Run("Writing a budget succeeds", func(t *testing.T) {
		newID, err := dao.WriteBudget(budget)
		assert.Nil(t, err)
		assert.Equal(t, newID, &id)
	})

	t.Run("Read an existing Budget", func(t *testing.T) {
		savedBudget, err := dao.ReadBudget(id)
		assert.Nil(t, err)
		assert.Equal(t, savedBudget, &budget)
	})

	t.Run("Delete an existing Budget", func(t *testing.T) {
		deletedID, err := dao.DeleteBudget(id)
		assert.Nil(t, err)
		assert.Equal(t, deletedID, &id)
	})

	t.Run("Read a deleted Budget", func(t *testing.T) {
		savedBudget, err := dao.ReadBudget(id)
		assert.Error(t, err, "sql: no rows in result set")
		assert.Nil(t, savedBudget)
	})
}

func TestStorageDao_AccountFunctions(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"

	account := Account{
		Name: name,
		ID:   id,
	}

	t.Run("Write a new Account", func(t *testing.T) {
		_, err := dao.WriteAccount(account)
		if err != nil {
			t.Fatal("Failed to write account to the MySQL DB!", err)
		}
	})

	t.Run("Read a new Account", func(t *testing.T) {
		savedAccount, err := dao.ReadAccount(id)
		if err != nil {
			t.Fatal("Failed to Read account from the MySQL DB!", err)
		}
		assert.Equal(t, savedAccount, &account)
	})

	t.Run("Delete a new Account", func(t *testing.T) {
		deletedAccount, err := dao.DeleteAccount(id)
		if err != nil {
			t.Fatal("Failed to delete transaction from the MySQL DB!", err)
		}

		assert.Equal(t, *deletedAccount, id)
	})

	t.Run("Read a deleted Transaction", func(t *testing.T) {
		deletedAccount, err := dao.ReadAccount(id)
		assert.Error(t, err, "sql: no rows in result set")
		assert.Nil(t, deletedAccount, nil)
	})
}

func TestStorageDao_WriteTransaction(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	amount := 100.00
	memo := "Test Memo"
	accountId := uuid.New()
	categoryId := uuid.New()
	transactionId := uuid.New()
	location, _ := time.LoadLocation("UTC")
	date := time.Date(2024, 12, 1, 4, 47, 10, 0, location)
	account := Account{
		Name: "DummyAccount",
		ID:   accountId,
	}
	category := Category{
		Title:          "DummyCategory",
		AllocatedFunds: 40.3,
		BudgetID:       uuid.UUID{},
		ID:             categoryId,
		Total:          10.4,
		Allocations:    []string{"test1", "test2"},
	}

	_, err = dao.WriteAccount(account)
	if err != nil {
		return
	}
	_, err = dao.WriteCategory(category)
	if err != nil {
		return
	}

	transaction := Transaction{
		Amount:   amount,
		Memo:     memo,
		Account:  account,
		Category: category,
		ID:       transactionId,
		Date:     date,
	}

	t.Run("Write a new Transaction", func(t *testing.T) {
		_, err := dao.WriteTransaction(transaction)
		if err != nil {
			t.Fatal("Failed to write Transaction to the MySQL DB!", err)
		}
	})

	t.Run("Read a new Transaction", func(t *testing.T) {
		savedTransaction, err := dao.ReadTransaction(transactionId)
		if err != nil {
			t.Fatal("Failed to Read Transaction from the MySQL DB!", err)
		}
		assert.Equal(t, &transaction, savedTransaction)
	})

	t.Run("Delete a new Transaction", func(t *testing.T) {
		deletedTransaction, err := dao.DeleteTransaction(transactionId)
		if err != nil {
			t.Fatal("Failed to delete transaction from the MySQL DB!", err)
		}

		assert.Equal(t, deletedTransaction, &transactionId)
	})

	t.Run("Read a deleted Transaction", func(t *testing.T) {
		deletedTransaction, err := dao.ReadTransaction(transactionId)
		assert.Error(t, err, "sql: no rows in result set")
		assert.Nil(t, deletedTransaction)
	})
}

func TestStorageDao_WriteCategory(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	title := "Test Title"
	allocatedFunds := 425.3
	total := 44.7
	allocations := []string{"allocation1", "allocation2"}

	category := Category{
		ID:             id,
		Title:          title,
		AllocatedFunds: allocatedFunds,
		Total:          total,
		Allocations:    allocations,
	}

	t.Run("Write a category", func(t *testing.T) {
		_, err := dao.WriteCategory(category)
		if err != nil {
			t.Fatal("Falied to write category to the MySQL DB!", err)
		}
	})

	t.Run("Read a category", func(t *testing.T) {
		savedCategory, err := dao.ReadCategory(id)
		if err != nil {
			t.Fatal("Failed to read category from the MySQL DB!", err)
		}
		assert.Equal(t, savedCategory, category)
	})

	t.Run("Delete a category", func(t *testing.T) {
		deletedCategory, err := dao.DeleteCategory(id)
		if err != nil {
			t.Fatal("Failed to delete transaction from the MySQL DB!", err)
		}

		assert.Equal(t, *deletedCategory, id)
	})

	t.Run("Read a deleted category", func(t *testing.T) {
		deletedCategory, err := dao.ReadBudget(id)
		assert.Error(t, err, "sql: no rows in result set")
		assert.Nil(t, deletedCategory, nil)
	})
}
