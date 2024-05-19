package persistance

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
	"time"
)

func Test_CreateTables(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	_ = dao.GormDB.Migrator().CreateTable(&Budget{})
	_ = dao.GormDB.Migrator().CreateTable(&Account{})
	_ = dao.GormDB.Migrator().CreateTable(&Allocation{})
	_ = dao.GormDB.Migrator().CreateTable(&Transaction{})
	_ = dao.GormDB.Migrator().CreateTable(&Category{})
}

func Test_AccountCRUDOperations(t *testing.T) {
	dao, err := NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	name := "test-account-003"

	account := Account{
		Name: name,
	}

	t.Run("Write a new Account", func(t *testing.T) {
		result := dao.GormDB.Create(&account)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Account", func(t *testing.T) {
		readAccount := Account{
			Model: gorm.Model{ID: account.ID},
		}
		result := dao.GormDB.Find(&readAccount)

		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
		assert.Equal(t, account.Name, readAccount.Name)
		assert.Equal(t, account.ID, readAccount.ID)
	})

	t.Run("Delete a new Account", func(t *testing.T) {
		deletedAccount := Account{
			Model: gorm.Model{ID: account.ID},
		}
		result := dao.GormDB.Delete(&deletedAccount)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a deleted Account", func(t *testing.T) {
		readAccount := Account{
			Model: gorm.Model{ID: account.ID},
		}
		result := dao.GormDB.Find(&readAccount)

		assert.NoError(t, result.Error)
		assert.Zero(t, result.RowsAffected)
	})
}

func Test_BudgetCRUDOperations(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	name := "test-account-001"

	budget := Budget{
		Name: name,
	}

	t.Run("Write a new Budget", func(t *testing.T) {
		result := dao.GormDB.Create(&budget)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Budget", func(t *testing.T) {
		readBudget := Budget{
			Model: gorm.Model{ID: budget.ID},
		}
		result := dao.GormDB.Find(&readBudget)

		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
		assert.Equal(t, readBudget.Name, budget.Name)
	})

	t.Run("Delete a new Budget", func(t *testing.T) {
		deletedBudget := Budget{
			Model: gorm.Model{ID: budget.ID},
		}
		result := dao.GormDB.Delete(&deletedBudget)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a deleted Budget", func(t *testing.T) {
		readBudget := Budget{
			Model: gorm.Model{ID: budget.ID},
		}
		result := dao.GormDB.Find(&readBudget)

		assert.NoError(t, result.Error)
		assert.Zero(t, result.RowsAffected)
	})
}

func Test_CategoryCRUDOperations(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	budgetName := "test-budget-001"

	budget := Budget{
		Name: budgetName,
	}
	dao.GormDB.Create(&budget)

	category := Category{
		Title:    "test-category",
		BudgetID: budget.ID,
		Total:    0,
	}

	t.Run("Write a new Category", func(t *testing.T) {
		result := dao.GormDB.Create(&category)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Category", func(t *testing.T) {
		readCategory := Category{
			Model: gorm.Model{ID: category.ID},
		}
		result := dao.GormDB.Find(&readCategory)

		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
		assert.Equal(t, readCategory.Title, category.Title)
	})

	t.Run("Read all categories for a budget", func(t *testing.T) {
		categories, err := dao.GetCategoriesForBudget(budget.ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, categories)
	})

	t.Run("Delete a new Category", func(t *testing.T) {
		deletedBudget := Category{
			Model: gorm.Model{ID: category.ID},
		}
		result := dao.GormDB.Delete(&deletedBudget)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a deleted Category", func(t *testing.T) {
		readBudget := Category{
			Model: gorm.Model{ID: category.ID},
		}
		result := dao.GormDB.Find(&readBudget)

		assert.NoError(t, result.Error)
		assert.Zero(t, result.RowsAffected)
	})
}

func Test_AllocationCRUDOperations(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	budgetName := "test-budget-001"
	budget := Budget{
		Name: budgetName,
	}
	dao.GormDB.Create(&budget)

	category := Category{
		Title:    "test-category",
		BudgetID: budget.ID,
		Total:    0,
	}
	dao.GormDB.Create(&category)

	allocation := Allocation{
		Amount:     47.0,
		CategoryID: category.ID,
	}

	t.Run("Write a new Allocation", func(t *testing.T) {
		result := dao.GormDB.Create(&allocation)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Allocation", func(t *testing.T) {
		readAllocation := Allocation{
			Model: gorm.Model{ID: allocation.ID},
		}
		result := dao.GormDB.Find(&readAllocation)

		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
		assert.Equal(t, readAllocation.Amount, allocation.Amount)
		assert.Equal(t, readAllocation.CategoryID, allocation.CategoryID)
	})

	t.Run("Delete a new Allocation", func(t *testing.T) {
		deletedAllocation := Allocation{
			Model: gorm.Model{ID: allocation.ID},
		}
		result := dao.GormDB.Delete(&deletedAllocation)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Allocation", func(t *testing.T) {
		readAllocation := Allocation{
			Model: gorm.Model{ID: allocation.ID},
		}
		result := dao.GormDB.Find(&readAllocation)

		assert.NoError(t, result.Error)
		assert.Zero(t, result.RowsAffected)
	})
}

func Test_TransactionCRUDOperations(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("PASSWORD"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	budgetName := "test-budget-001"
	budget := Budget{
		Name: budgetName,
	}
	dao.GormDB.Create(&budget)

	category := Category{
		Title:    "test-category",
		BudgetID: budget.ID,
		Total:    0,
	}
	dao.GormDB.Create(&category)

	allocation := Allocation{
		Amount:     47.0,
		CategoryID: category.ID,
	}
	dao.GormDB.Create(&allocation)

	name := "test-account-003"

	account := Account{
		Name: name,
	}
	dao.GormDB.Create(&account)

	transaction := Transaction{
		Amount:     23,
		Memo:       "Some Memo Test",
		AccountID:  account.ID,
		CategoryID: category.ID,
		Date:       time.Now(),
	}

	t.Run("Write a new Transaction", func(t *testing.T) {
		result := dao.GormDB.Create(&transaction)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Transaction", func(t *testing.T) {
		readTransaction := Transaction{
			Model: gorm.Model{ID: transaction.ID},
		}
		result := dao.GormDB.Find(&readTransaction)

		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
		assert.Equal(t, readTransaction.Amount, transaction.Amount)
		assert.Equal(t, readTransaction.CategoryID, transaction.CategoryID)
	})

	t.Run("Delete a new Transaction", func(t *testing.T) {
		deletedTransaction := Transaction{
			Model: gorm.Model{ID: transaction.ID},
		}
		result := dao.GormDB.Delete(&deletedTransaction)
		assert.NoError(t, result.Error)
		assert.NotZero(t, result.RowsAffected)
	})

	t.Run("Read a new Transaction", func(t *testing.T) {
		readTransaction := Transaction{
			Model: gorm.Model{ID: transaction.ID},
		}
		result := dao.GormDB.Find(&readTransaction)

		assert.NoError(t, result.Error)
		assert.Zero(t, result.RowsAffected)
	})
}
