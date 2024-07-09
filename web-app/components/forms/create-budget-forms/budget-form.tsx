"use client";

import React, {useState} from "react";
import {Budget, Category} from "@/app/types/BudgetTypes";
import CategoryForm from "@/components/forms/create-budget-forms/category-form";
// @ts-ignore
import {v4 as uuidv4} from 'uuid';

export interface BudgetProps {
    name: string,
    budget: Budget,
}

export default function BudgetForm(props: BudgetProps)  {
    const [formData, setFormData] = useState({
        name: props.name,
        budget: props.budget || {
            categories: [],
            name: "",
            ID: uuidv4(),
        },
    })

    const handleChange = (event: { target: { name: any; value: any; }; }) => {
        const { name, value } = event.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]:value,
        }))
    }

    const addCategory = (e: { preventDefault: () => void; }) => {
        setFormData((prevState) => ({
            ...prevState,
            budget:
                {
                    ...prevState.budget,
                    Categories: [
                        ...prevState.budget.Categories,
                        {
                            Total: 0,
                            BudgetID: 0,
                            Title: "Category Title",
                            Allocated: 0,
                            Spent: 0,
                            ID: uuidv4(),}
                    ]
                }
        }))
        e.preventDefault();
    }

    const handleCategoryChange = (category: Category) => {
        setFormData((prevState) => {
            console.log(prevState.budget.Categories)
            let categoryIndex = prevState.budget.Categories.findIndex(x => {
                return x.ID === category.ID
            });
            return {
            ...prevState,
            budget:
                {
                    ...prevState.budget,
                    Categories: [
                        ...prevState.budget.Categories.slice(0,categoryIndex),
                        category,
                        ...prevState.budget.Categories.slice(categoryIndex + 1),
                    ]
                }
        }})
    }

    return (
        <form>
            <label>
                Name:
                <input type="text" name="name" value={formData.name} onChange={handleChange}/>
            </label>
            <button onClick={addCategory}>Add Category</button>
            {
                formData.budget.Categories.map((category:Category) => (
                <CategoryForm category={category} key={category.ID} onChange={handleCategoryChange}></CategoryForm>
            ))}
            <input type="submit" value="Submit" />
        </form>
    );
}