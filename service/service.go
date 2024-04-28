package service

import (
	"BudgetingApp/persistance"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Start() error {
	dao, err := persistance.NewStorageDao(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), "budgetApp")
	if err != nil {
		return err
	}
	g := setupServer(dao)
	err = g.Run()
	return err
}

func setupServer(dao *persistance.StorageDao) *gin.Engine {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "Hello World!",
		})
	})

	g.GET("/budget", func(c *gin.Context) {
		budgets, err := getAllBudgets(dao)
		if err != nil {
			return
		}
		response := map[string][]Budget{
			"Budgets": budgets,
		}
		c.JSON(http.StatusOK, response)
	})

	g.GET("/budget/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(400, gin.H{"Message": "Invalid ID"})
			return
		}
		budget, err := dao.GetBudget(query.ID)
		if err != nil {
			c.JSON(400, gin.H{"Message": "No Budget found for ID"})
			return
		}
		c.JSON(200, toExternal(budget))
	})
	// Routes
	// /budget -> Get all budgets for user
	// /budget/${id} -> Get single budget with categories and computations
	// /account/${id} -> Get account and transactions
	//
	//

	return g
}

func getAllBudgets(dao *persistance.StorageDao) (budgets []Budget, err error) {
	internalBudgets, err := dao.GetAllBudgets()
	if err != nil {
		log.Fatal("Failed to read all budgets!", err)
		return
	}
	for _, budget := range internalBudgets {
		budgets = append(budgets, toExternal(budget))
	}
	log.Default().Println("Returning budgets: ", budgets)
	return
}

func getBudget(dao *persistance.StorageDao, id uint) (budget Budget, err error) {
	internalBudget, err := dao.GetBudget(id)
	if err != nil {
		log.Fatal("Failed to read budget!", err)
		return
	}
	budget = toExternal(internalBudget)

	log.Default().Println("Returning budget: ", budget)
	return
}

func toExternal(budget persistance.Budget) Budget {
	return Budget{
		Name: budget.Name,
		ID:   budget.ID,
	}
}

type Query struct {
	ID uint `uri:"name"`
}

type Budget struct {
	Name string
	ID   uint
}

type Category struct {
	Title       string
	Budget      Budget
	ID          uint
	Total       float64
	Allocations []Allocation
}

type Allocation struct {
	// date
	Amount     float64
	CategoryID string
	ID         uint
}

type Transaction struct {
	Amount   float64
	Memo     string
	Account  Account
	Category Category
	ID       uint
	// date
}

type Account struct {
	Name string
	ID   uint
}
