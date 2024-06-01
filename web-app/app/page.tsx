"use client";

import Link from "next/link";
import {useEffect} from "react";
import {getHelloWorld} from "@/components/api/client";

export default function Page() {
    useEffect(() => {
        console.log(getHelloWorld())
    }, [])

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
