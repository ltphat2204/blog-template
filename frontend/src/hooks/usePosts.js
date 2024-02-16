import Fetcher from '@/common/Fetcher'
import useSWR from 'swr'

export default function usePosts (offset, max) {
    const { data, error, isLoading } = useSWR(`/api/posts?offset=${offset}&max=${max}`, Fetcher)
    return {
      posts: data,
      isLoading,
      isError: error
    }
  }