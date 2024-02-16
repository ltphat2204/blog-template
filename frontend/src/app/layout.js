import { Montserrat } from "next/font/google";
import "./globals.css";
import Header from "@/components/Header";
import Footer from "@/components/Footer";

const montserrat = Montserrat({ subsets: ["latin", "vietnamese"] });

export const metadata = {
  title: "Blog title",
  description: "This is the blog of Some One, where they give their thoughts about something.",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${montserrat.className} p-4`}>
        <Header/>
        <main className=" max-w-container mx-auto">
          {children}
        </main>
        <Footer/>
      </body>
    </html>
  );
}
