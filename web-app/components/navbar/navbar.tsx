import Link from "next/link";
import './navbar.css';

export default function Navbar() {
    return (
        <main className="flex navbar">
            <ul className="horizontal">
                <li><Link href="/">Home</Link></li>
                <li><Link href="/">Budget</Link></li>
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
