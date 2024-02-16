import axios from 'axios'

export default async function Fetcher(url) {
    const res = await axios.get(url)
    return res.data
}