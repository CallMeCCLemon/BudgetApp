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

func Test_budget_CRUD_operations_tests(t *testing.T) {
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
