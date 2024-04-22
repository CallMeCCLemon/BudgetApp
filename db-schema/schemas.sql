CREATE DATABASE budgetApp;
Use budgetApp;

CREATE TABLE Accounts
(
    ID   CHAR(36) PRIMARY KEY,
    Name VARCHAR(255)
);

CREATE TABLE Budgets
(
    ID         CHAR(36) PRIMARY KEY,
    Name       VARCHAR(255),
    -- You can have separate tables for Categories and Accounts and then link them using foreign keys.
    -- But for simplicity, we'll just store comma-separated values for Categories and Accounts here.
    Categories TEXT,
    Accounts   TEXT
);

CREATE TABLE Categories
(
    ID             CHAR(36) PRIMARY KEY,
    Title          VARCHAR(255),
    AllocatedFunds DECIMAL(10, 2),
    BudgetID       CHAR(36),
    Total          DECIMAL(10, 2),
    Allocations    TEXT, -- Store UUIDs of Allocation
    FOREIGN KEY (BudgetID) REFERENCES Budgets (ID)
);

CREATE TABLE Allocations
(
    ID         CHAR(36) PRIMARY KEY,
    Amount     DECIMAL(10, 2),
    CategoryID CHAR(36),
    FOREIGN KEY (CategoryID) REFERENCES Categories (ID)
);

CREATE TABLE Transactions
(
    ID         CHAR(36) PRIMARY KEY,
    Amount     DECIMAL(10, 2),
    Memo       VARCHAR(255),
    AccountID  CHAR(36),
    CategoryID CHAR(36),
    Date       TIMESTAMP,
    FOREIGN KEY (AccountID) REFERENCES Accounts (ID),
    FOREIGN KEY (CategoryID) REFERENCES Categories (ID)
);


