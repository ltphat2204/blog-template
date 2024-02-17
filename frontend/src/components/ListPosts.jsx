import PostCard from "./PostCard";

export default function ListPosts({classname, posts}) {
    return (
        <div className={`${classname} flex flex-wrap w-full`}>
            {
                posts && posts.data.map(post => <PostCard data={post} key={post.id}/>)
            }
        </div>
    );
}