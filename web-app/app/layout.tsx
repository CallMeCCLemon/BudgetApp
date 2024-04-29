import Navbar from '../components/navbar/navbar'

export default function RootLayout(
    {
        children,
    }: {
        children: React.ReactNode;
    }) {
    return (
        <html lang="en">
            <Navbar />
            <main>{children}</main>
        </html>
    );
}
