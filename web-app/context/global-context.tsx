"use client";
import React, {useState} from "react";

interface Budget {
    Name: string;
    ID: number;
    Categories: Category[];
}

interface Category {
    "Title": string;
    "BudgetID": number;
    "ID": number;
    "Total": number;
    "Spent": number;
    "Allocated": number;
}

interface IGlobalContextProps {
    user: any;
    budget: Budget;
    loading: boolean;
    setUser: (user: any) => void;
    setLoading: (loading: boolean) => void;
    setBudget: (budget: Budget) => void;
}

const DEFAULT_BUDGET = {
    Name: "",
    ID: -1,
    Categories: [],
} as Budget

export const GlobalContext = React.createContext<IGlobalContextProps>({
    user: {},
    budget: DEFAULT_BUDGET,
    loading: true,
    setUser: () => {},
    setLoading: () => {},
    setBudget: () => {},
});

export const GlobalContextProvider = (props: any) => {
    const [currentUser, setCurrentUser] = useState({});
    const [isLoading, setIsLoading] = useState(true);
    const [currentBudget, setCurrentBudget] = useState(DEFAULT_BUDGET);

    return (
        <GlobalContext.Provider
            value={{
                user: currentUser,
                budget: currentBudget,
                loading: isLoading,
                setUser: setCurrentUser,
                setLoading: setIsLoading,
                setBudget: setCurrentBudget,
            }}
        >
            {props.children}
        </GlobalContext.Provider>
    );
};