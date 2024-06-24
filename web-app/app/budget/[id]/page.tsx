'use client'

import {Budget} from "@/app/types/BudgetTypes";
import dummyBudgetData from '../../../test/fixtures/budget.json'
import {useEffect, useState, useContext} from "react";
import {useParams} from "next/navigation";
import Table from "@/components/table/transaction-table";
import {GlobalContext} from "@/context/global-context";
import {COLUMN_TYPE} from "@/components/table/table";

export default function Page() {
    const params = useParams<{id: string}>();
    const [newCategory, setNewCategory] = useState<string>("")
    const [budgetData, setBudgetData] = useState<Budget | undefined>()
    const {budget, setBudget} = useContext(GlobalContext);

    useEffect(() => {
        // TODO: Load Budget Data here!
        const budgetId = params.id;
        if (!budget?.ID || budgetId != budget.ID.toString()) {
            // Call backend
            console.log("Call backend");
            setBudgetData(dummyBudgetData);
            setBudget(budgetData!);
        }
    }, [budget, budgetData, params, setBudget])

    const createNewCategory = () => {
        console.log(`Creating new Category! ${newCategory}`)
        setNewCategory("")
        // TODO: Call CreateNewCategory API.
    }

    const updateNewCategory = (event: React.FormEvent<HTMLInputElement>) => {
        setNewCategory(event.currentTarget.value)
    }

    const columns = [
        {name: "Title", type: COLUMN_TYPE.string},
        {name: "ID", type: COLUMN_TYPE.number},
        {name: "Total", type: COLUMN_TYPE.currency},
        {name: "Spent", type: COLUMN_TYPE.currency},
        {name: "Allocated", type: COLUMN_TYPE.currency},
    ]

    const getBudgetContent = (budgetData: Budget | undefined) => {
        if (budgetData !== undefined) {
            return (<div>
                <h1>{budgetData.Name}</h1>
                {newCategoryForm}
                <div className="w-full h-full flex flex-col">
                    <Table TableHeader={"Transactions"} columnNames={columns} rowData={budgetData.Categories} />
                </div>
            </div>)
        } else {
            // TODO: Implement Spinner here
            return <p>Spinner</p>
        }
    }

    const newCategoryForm =
        <div className="">
            <input placeholder="Category Name" value={newCategory} onInput={updateNewCategory}></input>
            <button onClick={createNewCategory}>Create</button>
        </div>

    return (
        <main className="flex min-h-screen flex-col p-6">
            {getBudgetContent(budgetData)}
        </main>
    );
}
