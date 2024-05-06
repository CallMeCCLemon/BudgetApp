'use client'

import { useSearchParams } from "next/navigation";

export default function Page() {
    const searchParams = useSearchParams();
    const budgetName = searchParams.get("name");

    return (
        <main className="flex min-h-screen flex-col p-6">
            <div>
                <p>Budget Page: {budgetName}</p>
            </div>
        </main>
    );
}
