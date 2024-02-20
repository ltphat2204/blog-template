import useSWR from 'swr'
import axios from 'axios';

export default function usePosts (page, category, search) {
    const { data, error, isLoading } = useSWR(`/api/posts?search=${search}&page=${page}&category_id=${category}`, async(url) => axios.get(url).then(res => res.data))
    return {
      posts: data,
      isPostsLoading: isLoading,
      isPostsError: error
    }
}