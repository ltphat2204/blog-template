'use client'
import Dropdown from "@/components/Dropdown";
import ListPosts from "@/components/ListPosts";
import Pagination from "@/components/Pagination";
import useCategories from "@/hooks/useCategories";
import usePosts from "@/hooks/usePosts";
import { useState, useEffect } from "react";

export default function Blogs() {
    const [search, setSearch] = useState('');
    const [searchTemp, setSearchTemp] = useState('');
    const [category, setCategory] = useState('');
    const [page, setPage] = useState(1);
    const { posts, isPostsLoading, isPostsError } = usePosts(page, category, search);
    const { categories, isCategoriesLoading } = useCategories();

    if (isPostsError) console.log(isPostsError);


    useEffect(() => {
        const handler = setTimeout(() => {
            setSearch(searchTemp);
        }, 500);

        return () => {
            clearTimeout(handler);
        };
    }, [searchTemp]);

    return (
        <div>
            <div className="flex md:mr-2 mb-4">
                <input className="w-full border-b border-gray-200 mr-4 py-1 px-2 focus:outline-none" placeholder="Type anything to search" onChange={e => setSearchTemp(e.target.value)}/>
                <Dropdown isLoading={isCategoriesLoading} options={categories ? categories.data : null} onSelect={setCategory} />
            </div>
            <ListPosts loading={isPostsLoading} posts={posts}/>
            <Pagination currentPage={page} onPageChange={setPage} totalPages={posts && posts.pagination ? Math.ceil(posts.pagination.total / 6) : 0}/>
        </div>
    )
}