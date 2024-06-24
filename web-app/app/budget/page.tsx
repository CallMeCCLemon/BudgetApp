"use client";
import {useContext, useEffect} from "react";
import {budgetService} from "@/services/budget-service";
import {GlobalContext} from "@/context/global-context";
import { Budget } from "../types/BudgetTypes";

export default function Page() {
    const {budgets, setBudgets} = useContext(GlobalContext);


    useEffect(() => {
        budgetService.getAllBudgets().then((budgets: Budget[]) => {
            setBudgets(budgets);
        })
    })

    return <div>
        {budgets.map(budget => {
            return <p key={budget.ID}>{budget.Name}</p>
        })}
    </div>
}