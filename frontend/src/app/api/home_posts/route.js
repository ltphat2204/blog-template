import axios from 'axios';
import { NextResponse } from 'next/server';

export async function GET() {
    try {
        const res = await axios.get(`${process.env.BACKEND_URI}/posts?page=1&limit=6`)
        const data = res.data
        return NextResponse.json(data)
    } catch (error) {
        console.error('Error:', error);
        return NextResponse.error(error.message, { status: 500 });
    }
    
}