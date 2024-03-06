'use client'
import axios from "axios";
import Image from "next/image";
import { useEffect, useState } from "react"
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import Markdown from 'react-markdown';
import '@/common/markdown.css'

export default function BlogDetail({params}) {
    const [post, setPost] = useState(null)

    useEffect(() => {
        axios.get(`/api/posts/${params.id}`)
            .then(res => setPost(res.data.data))
            .catch(err => console.log(err))
    }, [params])

    if (post === null) return (
        <div className="w-full h-full flex items-center justify-center">
            <AiOutlineLoading3Quarters className="animate-spin ml-2 text-2xl" />
        </div>
    )

    return (
        <div className="post">
            <div className="relative w-full h-72">
                <Image
                    className="object-cover object-center rounded"
                    fill
                    alt="Banner"
                    src={post.banner}
                />
            </div>
            <h1 className="mt-4 text-center">{post.title}</h1>
            <span className="text-gray-500 mt-2 text-justify">{post.description}</span>
            <Markdown>{post.content}</Markdown>
        </div>
    );
}