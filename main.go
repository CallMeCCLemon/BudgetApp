package main

import (
	"BudgetingApp/persistance"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
)

func main() {
	dao, err := persistance.NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		return
	}

	r := gin.Default()
	r.GET("/budget", func(c *gin.Context) {
		budgets, err := getAllBudgets(dao)
		response := map[string][]Budget{
			"Budgets": budgets,
		}
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, response)
	})
	// Routes
	// /budget -> Get all budgets for user
	// /budget/${id} -> Get single budget with categories and computations
	// /account/${id} -> Get account and transactions
	//
	//
	err = r.Run()
	if err != nil {
		return
	}
}

func getAllBudgets(dao *persistance.StorageDao) (budgets []Budget, err error) {
	internalBudgets, err := dao.ReadBudgets()
	if err != nil {
		return
	}
	for _, budget := range internalBudgets {
		budgets = append(budgets, toExternal(budget))
	}
	return
}

func toExternal(budget persistance.Budget) Budget {
	return Budget{
		Name: budget.Name,
		ID:   budget.ID,
	}
}

type Budget struct {
	Name string
	ID   uuid.UUID
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
	//internalBudget, err := storageDao.ReadBudget(id)
	//if err != nil {
	//	return Budget{}, err
	//}
	//var categories []Category
	//var errors []error
	//for _, categoryId := range internalBudget.Categories {
	//	category, err := storageDao.ReadCategory(categoryId)
	//	if err != nil {
	//		errors = append(errors, err)
	//	}
	//	categories = append(categories, Category{
	//		Title: category.Title,
	//		//Allocations:
	//	})
	//}
	budget = Budget{
		//Name: internalBudget.Name,
		//Categories: internalBudget.Categories,

	}

	return
}
