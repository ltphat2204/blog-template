'use client'
import Pagination from "@/components/Pagination";
import axios from "axios";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function Posts() {
    const [posts, setPosts] = useState(null);
    const [page, setPage] = useState(1);

    useEffect(()=>{
        axios.get('/api/posts', {
            params: {
                page
            }
        }).then(res => {
            setPosts(res.data)
        })
        .catch(err => console.log(err))
    }, [page]);

    return (
        <>
            <div className="w-full">
            {
                posts && posts.data.map((post, index) => (
                    <div className="flex p-4 bg-gray-200 mb-4 rounded w-full" key={index}>
                        <div className="flex">
                            <div className="flex-1/4">{post.title}</div>
                            <div className="short-description whitespace-normal flex-2/3">{post.description}</div>
                            <div className="flex-1/4">{post.category.name}</div>
                        </div>
                        <div className="whitespace-nowrap">
                            <Link href={`/dashboard/posts/edit/${post.id}`}>Edit</Link> | 
                            <Link href={`/dashboard/posts/remove/${post.id}`}> Remove</Link>
                        </div>
                    </div>
                ))
            }
            </div>
            <Pagination currentPage={page} onPageChange={setPage} totalPages={posts && posts.pagination ? Math.ceil(posts.pagination.total / 6) : 0}/>
        </>
    );
}