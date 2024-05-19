'use client'

import * as fs from 'fs';
import * as path from 'path';

import {useSearchParams} from "next/navigation";
import {Budget, Category} from "@/app/types/BudgetTypes";
import budgetData from '../../test/fixtures/budget.json'

export default function Page() {
    console.log(budgetData as Budget)
    const budget: Budget = budgetData

    const tableData = budgetData.Categories.map((category: Category, idx: number) => {
        return <tr key={category.ID}>
            <td>{category.Title}</td>
            <td>${category.Allocated}</td>
            <td>${category.Spent}</td>
            <td>${category.Total}</td>
        </tr>;
    });

    return (
        <main className="flex min-h-screen flex-col p-6">
            <div className="w-full flex-row mt-12">
                <h1>{budget.Name}</h1>
                <table className="table-auto w-full">
                    <thead>
                    <tr>
                        <th>Category</th>
                        <th>Allocated</th>
                        <th>Spent</th>
                        <th>Total</th>
                    </tr>
                    </thead>
                    <tbody>
                        {tableData}
                    </tbody>
                </table>
            </div>
        </main>
    );
}
