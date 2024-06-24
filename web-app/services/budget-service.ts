import api from "./api";
import {Budget, Category} from "@/app/types/BudgetTypes";

class BudgetService {
    async getAllBudgets(): Promise<Budget[]> {
        const response = await api.get("/budget");
        const budgets = response.data;
        return budgets.map((budget: any): Budget => this.convertToBudget(budget));
    }

    private convertToBudget(budget: any): Budget {
            return {
                Name: budget.Name,
                ID: budget.ID,
                Categories: budget.Categories.map((category: any) => this.convertToCategory(category)),
            }
    }

    private convertToCategory(category: any): Category {
        return {
            Title: category.Title,
            BudgetID: category.BudgetID,
            ID: category.ID,
            Total: category.Total,
            Spent: category.Spent,
            Allocated: category.Allocated,
        }
    }
}

export const budgetService = new BudgetService();
