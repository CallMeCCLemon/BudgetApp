package persistance

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestStorageDao_WriteBudget(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"
	categories := []uuid.UUID{uuid.New()}
	accounts := []uuid.UUID{uuid.New()}

	t.Run("Write then Read something", func(t *testing.T) {
		budget := Budget{
			Name:       name,
			ID:         id,
			Categories: categories,
			Accounts:   accounts,
		}
		_, err := dao.WriteBudget(budget)
		if err != nil {
			t.Fatal("Failed to write budget to the MySQL DB!", err)
		}
		//assert.Equal(t, new_id, id)
		savedAccount, err := dao.ReadBudget(id)
		if err != nil {
			t.Fatal("Failed to read budget from the MySQL DB!", err)
		}
		assert.Equal(t, savedAccount.ID, id)
		assert.Equal(t, savedAccount.Name, name)
	})
}

func TestStorageDao_WriteAccount(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"

	t.Run("Write then Read something", func(t *testing.T) {
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
