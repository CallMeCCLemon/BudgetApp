import Link from "next/link";
import './navbar.css';

export default function Navbar() {
    let budgets = ["Budget 1", "Budget 2", "Budget 3"]

    return (
        <main className="flex navbar">
            <ul className="horizontal">
                <li><Link href="/">Home</Link></li>
                <li className="dropdown">
                    <Link href="/budget">Budget</Link>
                    <div className="dropdown-content">
                        {budgets.map(budget => {
                            let link = '/budget?name=' + encodeURI(budget);
                            return <Link key={budget} href={link}>{budget}</Link>
                        })}
                    </div>
                </li>
                <li className="dropdown">
                    <Link href="javascript:void(0)" className="dropbtn">Account</Link>
                    <div className="dropdown-content right">
                    <Link href="./login">Login</Link>
                    <Link href="#">Create Account</Link>
                    </div>
                </li>
            </ul>
        </main>
    );
}
