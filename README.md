# Budgeting App

## 1.0 Features

* CRUD Budget - DONE
  * Create budget
  * Read budget
  * Delete budget
* CRUD Category
  * Create Category
  * Delete Category
  * Rename existing category
* CRUD Transaction
  * Create transaction manually
  * Delete transaction manually
  * Assign transaction to category
* CRUD Financial Account
* Upload CSV (or other common format) to the app. Parse all transactions and de-dupe.

## HTTP Data Types
* Budget
  * 

## Storage Data Types

* Budget
  * Name - string 
  * Categories - UUID[]
  * Accounts - UUID[]
  * ID - UUID
* Category
  * Title - String
  * Monthly Allocated Funds - Float 2 Decimals
  * BudgetID
  * Category Checkpoints - []CategoryCheckpoint
  * ID - UUID
* CategoryCheckpoint
  * Date Completed
  * Amount
  * CategoryID
  * ID
* Transaction
  * Amount - Float 2 decimals
  * Memo - String
  * Date - Date
  * AccountID
  * CategoryID
  * ID - UUID
* Financial Account
  * InstitutionName - String
  * ID - UUID

### Connecting to mysql

This will start the port-forwarding on your local machine to get started with mySql.
```
kubectl port-forward service/budget-app-mysql-database -n budget-app 3306
```
