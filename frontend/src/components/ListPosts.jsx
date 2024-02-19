import PostCard from "./PostCard";
import { AiOutlineLoading3Quarters } from "react-icons/ai";

export default function ListPosts({classname, posts}) {
    if (posts == undefined) return (
        <div className="w-full h-full flex items-center justify-center">
            <AiOutlineLoading3Quarters className="animate-spin ml-2 text-2xl"/>
        </div>
    )

    if (posts.data == undefined || posts.data.length === 0) return (
        <div className="w-full h-full flex items-center justify-center text-3xl">
            There is no post now, come later!
        </div>
    )

    return (
        <div className={`${classname} flex flex-wrap w-full`}>
            {
                posts.data.map(post => <PostCard data={post} key={post.id}/>)
            }
        </div>
    );
}