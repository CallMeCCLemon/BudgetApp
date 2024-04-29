package service

import (
	"BudgetingApp/persistance"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_budget_CRUD_operations_test(t *testing.T) {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return
	}
	ts := httptest.NewServer(setupServer(dao))
	defer ts.Close()

	client := &http.Client{}

	budgetName := "integ-test-budget-001"

	budget := Budget{
		Name: budgetName,
	}

	t.Run("Create a budget", func(t *testing.T) {
		body, err := json.Marshal(budget)
		if err != nil {
			return
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/budget", ts.URL),
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		println(buf.String())
		var newBudget Budget
		_ = json.Unmarshal(buf.Bytes(), &newBudget)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, budget.Name, newBudget.Name)
		budget = newBudget
	})

	t.Run("Read a single budget", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/budget/%d", ts.URL, budget.ID))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Delete a budget", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/budget/%d", ts.URL, budget.ID), nil)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		resp, err := client.Do(req)

		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Read all budgets", func(t *testing.T) {
		resp, err := http.Get(
			fmt.Sprintf("%s/budget", ts.URL),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func Test_account_CRUD_operations_test(t *testing.T) {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return
	}
	ts := httptest.NewServer(setupServer(dao))
	defer ts.Close()

	client := &http.Client{}

	budgetName := "integ-test-budget-001"
	accountName := "integ-test-account-001"

	budget := persistance.Budget{
		Name: budgetName,
	}
	dao.GormDB.Create(&budget)
	account := Account{
		Name:     accountName,
		BudgetID: budget.ID,
	}

	t.Run("Create an account", func(t *testing.T) {
		body, err := json.Marshal(account)
		if err != nil {
			return
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/account", ts.URL),
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		println(buf.String())
		var newAccount Account
		_ = json.Unmarshal(buf.Bytes(), &newAccount)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, account.Name, newAccount.Name)

		account = newAccount
	})

	t.Run("Read a single account", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/account/%d", ts.URL, account.ID))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Delete an account", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/account/%d", ts.URL, account.ID), nil)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		resp, err := client.Do(req)

		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func Test_category_CRUD_operations_test(t *testing.T) {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return
	}
	ts := httptest.NewServer(setupServer(dao))
	defer ts.Close()

	client := &http.Client{}

	budgetName := "integ-test-budget-001"
	accountName := "integ-test-account-001"
	categoryTitle := "integ-test-category-001"

	budget := persistance.Budget{
		Name: budgetName,
	}
	dao.GormDB.Create(&budget)
	account := persistance.Account{
		Name:     accountName,
		BudgetID: budget.ID,
	}
	dao.GormDB.Create(&account)
	category := Category{
		Title:    categoryTitle,
		BudgetID: budget.ID,
		Total:    478,
	}

	t.Run("Create a category", func(t *testing.T) {
		body, err := json.Marshal(category)
		if err != nil {
			return
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/category", ts.URL),
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		println(buf.String())
		var newCategory Category
		_ = json.Unmarshal(buf.Bytes(), &newCategory)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, category.Title, newCategory.Title)

		category = newCategory
	})

	t.Run("Read a single Category", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/category/%d", ts.URL, category.ID))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Delete a Category", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/category/%d", ts.URL, category.ID), nil)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		resp, err := client.Do(req)

		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
