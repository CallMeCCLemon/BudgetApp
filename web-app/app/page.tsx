import Link from "next/link";
import {ArrowRightIcon} from "@heroicons/react/24/outline";


export default function Page() {
    return (
        <main className="flex min-h-screen flex-col p-6">
            <div>
                <Link
                    href="/login"
                    className="flex items-center gap-5 self-start rounded-lg bg-blue-500 px-6 py-3 text-sm font-medium text-white transition-colors hover:bg-blue-400 md:text-base"
                >
                    <span>Log in</span> <ArrowRightIcon className="w-5 md:w-6" />
                </Link>
            </div>
        </main>
    );
}
