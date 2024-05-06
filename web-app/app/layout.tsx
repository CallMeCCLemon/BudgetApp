import Navbar from '../components/navbar/navbar'
import { Suspense } from 'react';

export default function RootLayout(
    {
        children,
    }: {
        children: React.ReactNode;
    }) {
    return (
        <html lang="en">
            <Navbar />
            <Suspense>
                <main>{children}</main>
            </Suspense>
        </html>
    );
}
