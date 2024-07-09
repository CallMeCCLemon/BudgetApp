import "./page.css";
import BudgetForm from "@/components/forms/create-budget-forms/budget-form";
import {Budget} from "@/app/types/BudgetTypes";
import {v4 as uuidv4} from 'uuid';

export default function Page() {
    let uuid = uuidv4();

    return (
        <div className="flex min-h-screen flex-col p-6">
            <div>
                <BudgetForm name={""} budget={{Name:"", ID:uuid, Categories:[]}}></BudgetForm>
            </div>
        </div>
    );
}
