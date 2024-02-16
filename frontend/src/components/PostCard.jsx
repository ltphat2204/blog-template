'use client'

import DateTimeFormat from "@/common/DateTimeFormat";
import Image from "next/image";
import Link from "next/link";

export default function PostCard({data}) {
    return (
        <Link href={`/blog/${data.id}`} className="flex-1/3 p-2 bg-gray-100 rounded mb-4 hover:drop-shadow-md text-base">
            <div className="w-full h-48 relative">
                <Image
                    src={data.banner}
                    alt="banner"
                    className="rounded object-cover object-center"
                    fill
                />
            </div>
            <h3 className="text-2xl font-semibold mt-4">{data.title}</h3>
            <div className="text-justify mt-2 short-paragraph">{data.description}</div>
            <div className="text-gray-500 mt-2">
                {
                    DateTimeFormat(data.created_at, {
                        year: "numeric",
                        month: "long",
                        day: "2-digit"
                    })
                }
            </div>
        </Link>
    );
}