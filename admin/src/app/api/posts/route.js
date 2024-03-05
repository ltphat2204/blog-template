import axios from 'axios';
import { NextResponse } from 'next/server';

export async function GET(request) {
    try {
        const query = request.nextUrl.searchParams;

        const queryString = query.toString();
        const url = `${process.env.BACKEND_URI}/posts?${queryString}&limit=6`;

        const res = await axios.get(url);
        const data = res.data;

        return NextResponse.json(data);
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}
