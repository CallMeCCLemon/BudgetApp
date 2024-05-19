
export interface Category {
    Total: number;
    BudgetID: number;
    Title: string;
    Allocated: number;
    Spent: number;
    ID: number;
}

export interface Budget {
    ID: number;
    Name: string;
    Categories: Category[];
}