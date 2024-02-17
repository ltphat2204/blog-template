import useSWR from 'swr'

export default function usePosts (page, limit) {
    const { data, error, isLoading } = useSWR(`/api/posts?page=${page}&limit=${limit}`, async(url) => {
      await axios.get(url)
    })
    return {
      posts: data,
      isLoading,
      isError: error
    }
  }