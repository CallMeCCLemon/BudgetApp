import Navbar from '../components/navbar/navbar'
import { Suspense } from 'react';
import {GlobalContextProvider} from "@/context/global-context";

export default function RootLayout(
    {
        children,
    }: {
        children: React.ReactNode;
    }) {
    return (
        <html lang="en">
            <body>
                <GlobalContextProvider>
                    <Navbar />
                    <Suspense>
                        <main>{children}</main>
                    </Suspense>
                </GlobalContextProvider>
            </body>
        </html>
    );
}
