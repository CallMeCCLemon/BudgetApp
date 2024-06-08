package service

import (
	"BudgetingApp/persistance"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Start() error {
	dao, err := persistance.NewStorageDao(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		"budgetApp")
	if err != nil {
		return err
	}
	g := setupServer(dao)
	err = g.Run()
	return err
}

func setupServer(dao *persistance.StorageDao) *gin.Engine {
	g := gin.Default()
	config := cors.Default()
	g.Use(config)

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "Hello World!",
		})
	})

	addBudgetRoutes(g, dao)
	addAccountRoutes(g, dao)
	addCategoryRoutes(g, dao)
	addTransactionRoutes(g, dao)
	addAllocationRoutes(g, dao)

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
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
			return
		}
		budget, err := dao.GetBudget(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Budget found for ID"})
			return
		}

		categories, err := dao.GetCategoriesForBudget(budget.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to get categories for budget"})
			return
		}

		c.JSON(http.StatusOK, toExternalBudget(*budget, categories))
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
		c.JSON(http.StatusOK, toExternalBudget(internalBudget, []persistance.Category{}))
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
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
			return
		}
		account, err := dao.GetAccount(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Account found for ID"})
			return
		}

		transactions, err := dao.GetTransactionsForAccount(account.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to get transactions for Account"})
			return
		}

		c.JSON(http.StatusOK, toExternalAccount(*account, transactions))
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
		c.JSON(http.StatusOK, toExternalAccount(internalAccount, []persistance.Transaction{}))
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

func addCategoryRoutes(g *gin.Engine, dao *persistance.StorageDao) {
	g.GET("/category/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
			return
		}
		category, err := dao.GetCategory(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Category found for ID"})
			return
		}
		c.JSON(http.StatusOK, toExternalCategory(*category))
		return
	})

	g.POST("/category", func(c *gin.Context) {
		var category Category
		err := c.Bind(&category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Category"})
			return
		}
		internalCategory := toInternalCategory(category)
		result := dao.GormDB.Create(&internalCategory)
		if result.Error != nil || result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to create Category!"})
			return
		}
		c.JSON(http.StatusOK, toExternalCategory(internalCategory))
		return
	})

	g.DELETE("/category/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}

		result := dao.GormDB.Delete(&persistance.Category{}, query.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Category found for ID"})
			return
		}

		c.JSON(http.StatusOK, nil)
		return
	})

	return
}

func addTransactionRoutes(g *gin.Engine, dao *persistance.StorageDao) {
	g.GET("/transactions", func(c *gin.Context) {
		transactions, err := getAllTransactions(dao)
		if err != nil {
			return
		}
		response := map[string][]Transaction{
			"Transactions": transactions,
		}
		c.JSON(http.StatusOK, response)
		return
	})

	g.GET("/transaction/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
			return
		}
		transaction, err := dao.GetTransaction(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Category found for ID"})
			return
		}
		c.JSON(http.StatusOK, toExternalTransaction(*transaction))
		return
	})

	g.POST("/transaction", func(c *gin.Context) {
		var transaction Transaction
		err := c.Bind(&transaction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Transaction"})
			return
		}
		internalTransaction := toInternalTransaction(transaction)
		result := dao.GormDB.Create(&internalTransaction)
		if result.Error != nil || result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to create Transaction!"})
			return
		}
		c.JSON(http.StatusOK, toExternalTransaction(internalTransaction))
		return
	})

	g.DELETE("/transaction/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}

		result := dao.GormDB.Delete(&persistance.Transaction{}, query.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Transaction found for ID"})
			return
		}

		c.JSON(http.StatusOK, nil)
		return
	})

	return
}

func addAllocationRoutes(g *gin.Engine, dao *persistance.StorageDao) {
	g.GET("/allocation/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
			return
		}
		allocation, err := dao.GetAllocation(query.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Allocation found for ID"})
			return
		}
		c.JSON(http.StatusOK, toExternalAllocation(*allocation))
		return
	})

	g.POST("/allocation", func(c *gin.Context) {
		var allocation Allocation
		err := c.Bind(&allocation)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Transaction"})
			return
		}
		internalAllocation := toInternalAllocation(allocation)
		result := dao.GormDB.Create(&internalAllocation)
		if result.Error != nil || result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to create Allocation!"})
			return
		}
		c.JSON(http.StatusOK, toExternalAllocation(internalAllocation))
		return
	})

	g.DELETE("/allocation/:id", func(c *gin.Context) {
		var query Query
		err := c.ShouldBindUri(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}

		result := dao.GormDB.Delete(&persistance.Allocation{}, query.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No Allocation found for ID"})
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
		budgets = append(budgets, toExternalBudget(budget, []persistance.Category{}))
	}
	log.Default().Println("Returning budgets: ", budgets)
	return
}

func getAllTransactions(dao *persistance.StorageDao) (transactions []Transaction, err error) {
	internalTransactions, err := dao.GetAllTransactions()
	if err != nil {
		log.Fatal("Failed to read all budgets!", err)
		return
	}
	for _, transaction := range internalTransactions {
		transactions = append(transactions, toExternalTransaction(transaction))
	}
	log.Default().Println("Returning budgets: ", transactions)
	return
}
