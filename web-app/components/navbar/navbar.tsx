import Link from "next/link";
import './navbar.css';
import budgetData from '../../test/fixtures/budgets.json'
import dummyAccountData from '../../test/fixtures/accounts.json'
import {Budget} from "@/app/types/BudgetTypes";

export default function Navbar() {
    let budgets: Budget[] = budgetData.Budgets;
    let accounts = dummyAccountData.Accounts;

    return (
        <main className="flex navbar">
            <ul className="horizontal navbar-ul">
                <li className="navbar-li"><Link href="/">Home</Link></li>
                <li className="dropdown navbar-li">
                    <Link href={"/budget"}>Budgets</Link>
                    <div className="dropdown-content">
                        {budgets.map(budget => {
                            let link = `/budget/${budget.ID}`;
                            return <Link key={budget.ID} href={link}>{budget.Name}</Link>
                        })}
                    </div>
                </li>
                <li className="dropdown navbar-li">
                    <Link href={"/account"}>Accounts</Link>
                    <div className="dropdown-content">
                        {accounts.map(account => {
                            let link = `/account/${account.ID}`;
                            return <Link key={account.ID} href={link}>{account.Name}</Link>
                        })}
                    </div>
                </li>
                <li className="dropdown navbar-li">
                    <Link href="#" className="dropbtn navbar-a">Account</Link>
                    <div className="dropdown-content right">
                        <Link href="./login">Login</Link>
                        <Link href="./create-account">Create Account</Link>
                    </div>
                </li>
            </ul>
        </main>
    );
}
