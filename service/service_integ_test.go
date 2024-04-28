package service

import (
	"BudgetingApp/persistance"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_get_budgets_api_integration_tests(t *testing.T) {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return
	}
	ts := httptest.NewServer(setupServer(dao))
	defer ts.Close()

	t.Run("it should return all budgets which exist", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/budget", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		println(buf.String())
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})
}

func Test_post_budget_api_integration_tests(t *testing.T) {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return
	}
	ts := httptest.NewServer(setupServer(dao))
	defer ts.Close()

	t.Run("it should create a new budget", func(t *testing.T) {
		body := []byte(`{"name": "integ-test-budget-001"}`)

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
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})
}
