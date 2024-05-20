'use client'


import {Budget, Category} from "@/app/types/BudgetTypes";
import dummyBudgetData from '../../../test/fixtures/budget.json'
import {useEffect, useState} from "react";
import {usePathname, useRouter} from "next/navigation";

export default function Page() {
    const pathName = usePathname();
    const [newCategory, setNewCategory] = useState<string>("")
    const [budgetData, setBudgetData] = useState<Budget | undefined>()

    useEffect(() => {
        // TODO: Load Budget Data here!
        const splitPath = pathName.split('/')
        const budgetId = splitPath[splitPath.length-1]
        console.log(`Budget ID: ${budgetId}`)
        setBudgetData(dummyBudgetData);
    }, [pathName])

    const createNewCategory = () => {
        console.log(`Creating new Category! ${newCategory}`)
        setNewCategory("")
        // TODO: Call CreateNewCategory API.
    }

    const updateNewCategory = (event: React.FormEvent<HTMLInputElement>) => {
        setNewCategory(event.currentTarget.value)
    }

    const getBudgetContent = (budgetData: Budget | undefined) => {
        const tableData = budgetData?.Categories?.map((category: Category) => {
            return <tr key={category.ID}>
                <td>{category.Title}</td>
                <td>Pending Implementation</td>
                <td>${category.Allocated}</td>
                <td>${category.Spent}</td>
                <td>${category.Total}</td>
            </tr>;
        });

        if (budgetData !== undefined) {
            return (<div>
                <h1>{budgetData.Name}</h1>
            {
                newCategoryForm
            }
            <table className="table-auto w-full">
                <thead>
                <tr>
                    <th>Category</th>
                    <th>Rollover</th>
                    <th>Allocated</th>
                    <th>Spent</th>
                    <th>Total</th>
                </tr>
                </thead>
                <tbody>
                {tableData}
                </tbody>
            </table>
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
            <div className="w-full flex-row mt-12">
                {getBudgetContent(budgetData)}
            </div>
        </main>
    );
}
