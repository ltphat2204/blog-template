'use client'

import Link from "next/link";
import { useState } from "react";
import { FaFacebook, FaYoutube, FaInstagramSquare, FaSearch  } from "react-icons/fa";
import { GiHamburgerMenu } from "react-icons/gi";
import { ImCross } from "react-icons/im";

export default function Header() {
    const [toggleMobileOpen, setToggleMobileOpen] = useState(false);

    const toggleTransition = {
        "open": "translate-x-0 duration-500",
        "close": "translate-x-full duration-500"
    }

    return (
        <header className="p-4 pt-2 flex justify-between items-center">
            <div className="text-2xl flex">
                <FaFacebook className="mr-2" />
                <FaInstagramSquare className="mr-2" />
                <FaYoutube className="" />
            </div>
            <nav className="text-xl hidden md:block">
                <Link className="mr-4" href="/">Home</Link>
                <Link className="mr-4" href="/">Blogs</Link>
                <Link className="mr-4" href="/about">About</Link>
                <Link className="" href="/">Contact</Link>
            </nav>
            <div className="text-2xl flex">
                <FaSearch className="cursor-pointer"/>
                <GiHamburgerMenu className="ml-4 md:hidden" onClick={() => setToggleMobileOpen(true)}/>
            </div>
            <div className={`md:hidden w-screen h-screen flex justify-end fixed top-0 left-0 z-10 bg-gray-900/70 ${toggleMobileOpen ? toggleTransition.open : toggleTransition.close}`}>
                <div className="h-full w-72 bg-white p-8"> 
                    <div className="text-2xl w-full flex justify-end">
                        <ImCross onClick={() => setToggleMobileOpen(false)}/>
                    </div>
                    <nav className="text-xl flex flex-col mt-4">
                        <Link className="mb-4 py-2 border-b border-gray-500" href="/">Home</Link>
                        <Link className="mb-4 py-2 border-b border-gray-500" href="/">Blogs</Link>
                        <Link className="mb-4 py-2 border-b border-gray-500" href="/about">About</Link>
                        <Link className="py-2" href="/">Contact</Link>
                    </nav>
                </div>
            </div>
        </header>
    );
}