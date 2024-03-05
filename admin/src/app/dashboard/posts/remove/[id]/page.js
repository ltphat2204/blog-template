'use client'

import { useRouter } from 'next/navigation';
import axios from "axios";

export default function Remove({ params }) {
    const router = useRouter();
    const handleRemove = (e) => {
        e.preventDefault();
        axios.delete(`/api/posts/${params.id}`)
        .then(res => {
            if (res.data.success) {
                alert("Post deleted sucessful!")
                router.push('/dashboard/posts')
            }else {
                console.log(res.data.error)
            }
        })
        .catch(err => console.log(err))
    }

    return (
        <form onSubmit={handleRemove}>
            <h1>You are about to delete the blog with id: {params.id}. Are you sure? </h1>
            <div className="mt-4">
                <button className="px-4 py-2 bg-gray-700 text-gray-50 rounded" type="submit">Confirm</button>
                <button className="px-4 py-2 bg-gray-200 ml-4 rounded" type='button' onClick={()=>{router.push('/dashboard/posts')}}>Cancel</button>
            </div>
        </form>
    );
}