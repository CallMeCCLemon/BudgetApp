import Link from "next/link";
import './navbar.css';
import budgetData from '../../test/fixtures/budgets.json'
import {Budget} from "@/app/types/BudgetTypes";

export default function Navbar() {
    let budgets: Budget[] = budgetData.Budgets;

    return (
        <main className="flex navbar">
            <ul className="horizontal">
                <li><Link href="/">Home</Link></li>
                <li className="dropdown">
                    <Link href="#">Budgets</Link>
                    <div className="dropdown-content">
                        {budgets.map(budget => {
                            let link = `/budget/${budget.ID}`;
                            return <Link key={budget.ID} href={link}>{budget.Name}</Link>
                        })}
                    </div>
                </li>
                <li className="dropdown">
                    <Link href="#" className="dropbtn">Account</Link>
                    <div className="dropdown-content right">
                    <Link href="./login">Login</Link>
                    <Link href="#">Create Account</Link>
                    </div>
                </li>
            </ul>
        </main>
    );
}
