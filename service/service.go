package service

import (
	"BudgetingApp/persistance"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	addBudgetRoutes(g, dao)
	addAccountRoutes(g, dao)

	return g
}

func addBudgetRoutes(g *gin.Engine, dao *persistance.StorageDao) {
	g.GET("/budget", func(c *gin.Context) {
		budgets, err := getAllBudgets(dao)
		if err != nil {
			return
		}
		response := map[string][]Budget{
			"Budgets": budgets,
		}
		c.JSON(http.StatusOK, response)
		return
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
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Budget found for ID"})
			return
		}
		c.JSON(http.StatusOK, toExternalBudget(*budget))
		return
	})

	g.POST("/budget", func(c *gin.Context) {
		var budget Budget
		err := c.Bind(&budget)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid budget"})
			return
		}
		internalBudget := toInternalBudget(budget)
		result := dao.GormDB.Create(&internalBudget)
		if result.Error != nil || result.RowsAffected == 0 {
			c.JSON(500, gin.H{"Message": "Failed to create budget!"})
			return
		}
		c.JSON(http.StatusOK, toExternalBudget(internalBudget))
		return
	})

	g.DELETE("/budget/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}

		result := dao.GormDB.Delete(&persistance.Budget{}, query.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Budget found for ID"})
			return
		}

		c.JSON(http.StatusOK, nil)
		return
	})

	return
}

func addAccountRoutes(g *gin.Engine, dao *persistance.StorageDao) {
	g.GET("/account/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(400, gin.H{"Message": "Invalid ID"})
			return
		}
		budget, err := dao.GetBudget(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Account found for ID"})
			return
		}
		c.JSON(http.StatusOK, toExternalBudget(*budget))
		return
	})

	g.POST("/account", func(c *gin.Context) {
		var account Account
		err := c.Bind(&account)
		if err != nil {
			c.JSON(400, gin.H{"Message": "Invalid Account"})
			return
		}
		internalAccount := toInternalAccount(account)
		result := dao.GormDB.Create(&internalAccount)
		if result.Error != nil || result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to create Account!"})
			return
		}
		c.JSON(http.StatusOK, toExternalAccount(internalAccount))
		return
	})

	g.DELETE("/account/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}

		result := dao.GormDB.Delete(&persistance.Account{}, query.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Account found for ID"})
			return
		}

		c.JSON(http.StatusOK, nil)
		return
	})

	return
}

func getAllBudgets(dao *persistance.StorageDao) (budgets []Budget, err error) {
	internalBudgets, err := dao.GetAllBudgets()
	if err != nil {
		log.Fatal("Failed to read all budgets!", err)
		return
	}
	for _, budget := range internalBudgets {
		budgets = append(budgets, toExternalBudget(budget))
	}
	log.Default().Println("Returning budgets: ", budgets)
	return
}

func toExternalBudget(budget persistance.Budget) Budget {
	return Budget{
		Name: budget.Name,
		ID:   budget.ID,
	}
}

func toInternalBudget(budget Budget) persistance.Budget {
	return persistance.Budget{
		Name:  budget.Name,
		Model: gorm.Model{ID: budget.ID},
	}
}

func toExternalAccount(account persistance.Account) Account {
	return Account{
		ID:       account.ID,
		Name:     account.Name,
		BudgetID: account.BudgetID,
	}
}

func toInternalAccount(account Account) persistance.Account {
	return persistance.Account{
		Model:    gorm.Model{ID: account.ID},
		Name:     account.Name,
		BudgetID: account.BudgetID,
	}
}

type Query struct {
	ID uint `uri:"id"`
}

type Budget struct {
	Name string `uri:"name"`
	ID   uint   `uri:"id"`
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
	Name     string
	ID       uint
	BudgetID uint
}
