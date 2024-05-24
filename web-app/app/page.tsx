import Link from "next/link";
import {ArrowRightIcon} from "@heroicons/react/24/outline";


export default function Page() {
    return (
        <main className="flex min-h-screen flex-col p-6">
            <div>
                <table>
                    <tbody>
                        <tr>Create New Budget</tr>
                        <tr><Link href={"/budget"}>Budgets</Link></tr>
                        <tr><Link href={"/account"}>Accounts</Link></tr>
                    </tbody>
                </table>

            </div>
        </main>
    );
}
