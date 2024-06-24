"use client";
import React, {useState} from "react";
import {Category, Budget} from "@/app/types/BudgetTypes";

export interface IGlobalContextProps {
    user: any;
    budget: Budget;
    budgets: Budget[];
    loading: boolean;
    setUser: (user: any) => void;
    setLoading: (loading: boolean) => void;
    setBudget: (budget: Budget) => void;
    setBudgets: (budgets: Budget[]) => void;
}

const DEFAULT_BUDGET = {
    Name: "",
    ID: -1,
    Categories: [],
} as Budget

export const GlobalContext = React.createContext<IGlobalContextProps>({
    user: {},
    budget: DEFAULT_BUDGET,
    budgets: [DEFAULT_BUDGET],
    loading: true,
    setUser: () => {},
    setLoading: () => {},
    setBudget: () => {},
    setBudgets: () => {},
});

export const GlobalContextProvider = (props: any) => {
    const [currentUser, setCurrentUser] = useState({});
    const [isLoading, setIsLoading] = useState(true);
    const [currentBudget, setCurrentBudget] = useState(DEFAULT_BUDGET);
    const [budgets, setBudgets] = useState([DEFAULT_BUDGET])

    return (
        <GlobalContext.Provider
            value={{
                user: currentUser,
                budget: currentBudget,
                budgets: budgets,
                loading: isLoading,
                setUser: setCurrentUser,
                setLoading: setIsLoading,
                setBudget: setCurrentBudget,
                setBudgets: setBudgets,
            }}
        >
            {props.children}
        </GlobalContext.Provider>
    );
};