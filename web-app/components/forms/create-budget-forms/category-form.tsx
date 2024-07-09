import React, {useEffect, useState} from "react";
import {Category} from "@/app/types/BudgetTypes";

export default function CategoryForm(props: { category: Category, onChange?: (category: Category) => void })  {
    const [formData, setFormData] = useState({
        Total: props.category.Total,
        BudgetId: props.category.BudgetID,
        Title: props.category.Title,
        Allocated: props.category.Allocated,
        Spent: props.category.Spent,
        ID: props.category.ID,
    })

    useEffect(() => {
        if (props?.onChange && formData as unknown as Category != props.category) {
            props?.onChange(formData as unknown as Category);
        }
    })

    const handleChange = (event: { target: { name: any; value: any; }, preventDefault: () => void }) => {
        const { name, value } = event.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]:value,
        }))
        event.preventDefault();
    }

    return (<div>
                <label>
                    Category Name:
                    <input type="text" name="Title" value={props.category.Title} onChange={handleChange} />
                </label>
                <label>
                    Category Allocated:
                    <input type="number" name="Allocated" value={props.category.Allocated} onChange={handleChange} />
                </label>
                <label>
                    Category Spent:
                    <input type="number" name="Spent" value={props.category.Spent} onChange={handleChange} />
                </label>
                <label>
                    Category Total:
                    <input type="number" name="Total" value={props.category.Total} onChange={handleChange} />
                </label>
            </div>
    );
}