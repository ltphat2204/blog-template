import Link from "next/link";

export default function Footer() {
    return (
        <footer className="md:h-64 h-fit bg-gray-600 rounded text-white flex flex-col md:flex-row p-4 md:items-center mt-8">
            <div className="flex-1 py-2 md:px-2">
                <h5 className="text-3xl font-bold mb-2">End of blog</h5>
                <div>This is the brief information about blog website. This website created by Nextjs.</div>
            </div>
            <div className="flex-1 py-2 md:px-2">
                <h5 className="text-3xl font-bold mb-2">Basic information</h5>
                <ul>
                    <li><b>Adress:</b> Somewhere in somewhere at somewhere</li>
                    <li><b>Phone:</b> (028) 366 - 2654</li>
                    <li><b>Email:</b> example@gmail.com</li>
                </ul>
            </div>
            <div className="flex-1 py-2 md:px-2">
                <h5 className="text-3xl font-bold mb-2">Social media</h5>
                <ul>
                    <li>Facebook: <Link className="underline" href="/">My Facebook</Link></li>
                    <li>Instagram: <Link className="underline" href="/">My Instagram</Link></li>
                    <li>Youtube: <Link className="underline" href="/">My Youtube</Link></li>
                </ul>
            </div>
        </footer>
    );
}