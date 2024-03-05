import Link from "next/link";

export default function SideTab() {
    return (
        <nav className="min-h-screen bg-gray-800 w-64 flex-sidetab px-8 py-4 text-gray-100 text-lg">
            <h1 className="text-2xl font-semibold mb-8">ADMIN PAGE</h1>
            <Link className="pb-2 my-4 block border-b border-gray-100" href="/dashboard/accounts">Accounts</Link>
            <Link className="pb-2 my-4 block border-b border-gray-100" href="/dashboard/posts">Posts</Link>
        </nav>
    );
}