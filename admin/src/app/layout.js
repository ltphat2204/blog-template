'use client'

import { Montserrat } from "next/font/google";
import "./globals.css";
import {UserContextProvider} from "@/contexts/UserContext";

const montserrat = Montserrat({ subsets: ["latin", "vietnamese"] });

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${montserrat.className} p-4`}>
        <main className="h-screen">
          <UserContextProvider>
            {children}
          </UserContextProvider>
        </main>
      </body>
    </html>
  );
}
