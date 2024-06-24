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
kubectl port-forward service/budget-app-psql-rw -n budget-app 5432
```

### Create local kind server

```bash
kind create cluster --name budget-app
kind get kubeconfig --name budget-app > .kube/config

helm repo add cnpg https://cloudnative-pg.github.io/charts
helm upgrade --install cnpg \
  --namespace cnpg-system \
  --create-namespace \
  cnpg/cloudnative-pg
  
kubectl create secret generic mysql-passwords -n budget-app --from-literal=rootUser=root --from-literal=rootHost=% --from-literal=rootPassword="dummy-password123!"
```

### Run Docker container locally
```bash

# use docker image ls to get this value
IMAGE_ID=
PORT=5432 # default for psql
PSWD='somepasswd' # k8s secret.

docker run --rm -p 8080:8080 -it -p 5432:5432 \
  -e HOST='127.0.0.1' \
  -e PORT=5432 \
  -e USERNAME='budgetapp' \
  -e PASSWORD=$PSWD \
  $IMAGE_ID
```