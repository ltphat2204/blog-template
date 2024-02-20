import useSWR from 'swr'
import axios from 'axios';

export default function useCategories () {
    const { data, error, isLoading } = useSWR(`/api/categories`, async(url) => axios.get(url).then(res => res.data))
    return {
      categories: data,
      isCategoriesLoading: isLoading,
      isCategoriesError: error
    }
}