'use client'

import axios from "axios";
import Image from "next/image";
import { useEffect, useState, useContext } from "react"
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import { useRouter } from 'next/navigation';
import { UserContext } from '@/contexts/UserContext';

export default function Page({ params }) {
    const router = useRouter();
    const [post, setPost] = useState(null)
    const { user } = useContext(UserContext); 

    useEffect(() => {
        axios.get(`/api/posts/${params.id}`)
            .then(res => setPost(res.data.data))
            .catch(err => console.log(err))
    }, [params])

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setPost(prevPost => ({
            ...prevPost,
            [name]: value
        }));
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.patch(`/api/posts/${params.id}`, post, {
            headers: {
                Authorization: `BEARER ${user}`
            }
        })
        .then(res => {
            if (res.data.success) {
                alert("Post updated sucessful!")
                router.push('/dashboard/posts')
            }else {
                alert(res.data.message)
            }
        })
        .catch(err => console.log(err))
    }

    if (post === null) return (
        <div className="w-full h-full flex items-center justify-center">
            <AiOutlineLoading3Quarters className="animate-spin ml-2 text-2xl" />
        </div>
    )

    return (
        <form className="h-max" onSubmit={handleSubmit}>
            <div className="w-full h-72 relative mb-4">
                <Image
                    className="object-cover object-center rounded"
                    fill
                    alt="Banner"
                    src={post.banner}
                />
            </div>
            <div>
                <label htmlFor="banner" className="text-xl">Banner link: </label>
                <input
                    id="banner"
                    name="banner"
                    className="text-lg mb-2 focus:outline-none border-b border-gray-400 w-full"
                    value={post.banner}
                    onChange={handleInputChange}
                />
            </div>
            <div>
                <label htmlFor="title" className="text-xl">Title: </label>
                <input
                    id="title"
                    name="title"
                    className="text-lg mb-2 focus:outline-none border-b border-gray-400 w-full"
                    value={post.title}
                    onChange={handleInputChange}
                />
            </div>
            <div>
                <label htmlFor="description" className="text-xl">Description: </label>
                <textarea
                    id="description"
                    name="description"
                    className="text-lg mb-2 h-24 focus:outline-none border-b border-gray-400 w-full"
                    value={post.description}
                    onChange={handleInputChange}
                />
            </div>
            <div>
                <label htmlFor="content" className="text-xl">Content: </label>
                <textarea
                    id="content"
                    name="content"
                    className="text-lg h-72 focus:outline-none border-b border-gray-400 w-full"
                    value={post.content}
                    onChange={handleInputChange}
                />
            </div>
            <div className="mt-4">
                <button className="px-4 py-2 bg-gray-700 text-gray-50 rounded" type="submit">Confirm</button>
                <button className="px-4 py-2 bg-gray-200 ml-4 rounded" onClick={()=>{router.push('/dashboard/posts')}}>Cancel</button>
            </div>
        </form>
    );
}
