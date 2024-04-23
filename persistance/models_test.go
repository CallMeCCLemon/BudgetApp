package persistance

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"gotest.tools/v3/assert"
	"time"
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
		if err != nil {
			t.Fatal("Failed to write budget to the MySQL DB!", err)
		}
		assert.DeepEqual(t, newID, &id)
	})

	t.Run("Read an existing Budget", func(t *testing.T) {
		savedBudget, err := dao.ReadBudget(id)
		if err != nil {
			t.Fatal("Failed to read budget from the MySQL DB!", err)
		}
		assert.DeepEqual(t, savedBudget, budget)
	})
}

func TestStorageDao_WriteAccount(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"

	t.Run("Write then Read a new Account", func(t *testing.T) {
		account := Account{
			Name: name,
			ID:   id,
		}
		_, err := dao.WriteAccount(account)
		if err != nil {
			t.Fatal("Failed to write account to the MySQL DB!", err)
		}
		//assert.Equal(t, new_id, id)
		savedAccount, err := dao.ReadAccount(id)
		if err != nil {
			t.Fatal("Failed to Read account from the MySQL DB!", err)
		}
		assert.Equal(t, savedAccount.ID, id)
		assert.Equal(t, savedAccount.Name, name)
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

	t.Run("Write then Read a new Transaction", func(t *testing.T) {
		transaction := Transaction{
			Amount:   amount,
			Memo:     memo,
			Account:  account,
			Category: category,
			ID:       transactionId,
			Date:     date,
		}
		id, err := dao.WriteTransaction(transaction)
		if err != nil {
			t.Fatal("Failed to write Transaction to the MySQL DB!", err)
		}
		//assert.Equal(t, new_id, id)
		savedTransaction, err := dao.ReadTransaction(id)
		if err != nil {
			t.Fatal("Failed to Read Transaction from the MySQL DB!", err)
		}
		assert.DeepEqual(t, transaction, savedTransaction)
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

	t.Run("Write then read a category", func(t *testing.T) {
		category := Category{
			ID:             id,
			Title:          title,
			AllocatedFunds: allocatedFunds,
			Total:          total,
			Allocations:    allocations,
		}
		_, err := dao.WriteCategory(category)
		if err != nil {
			t.Fatal("Falied to write category to the MySQL DB!", err)
		}

		savedCategory, err := dao.ReadCategory(id)
		if err != nil {
			t.Fatal("Failed to read category from the MySQL DB!", err)
		}
		assert.DeepEqual(t, savedCategory, category)
	})
}
