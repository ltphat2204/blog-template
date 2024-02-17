import axios from 'axios';
import { NextResponse } from 'next/server';

export async function POST(req) {
    try {
        const newSubscriber = await req.json();
        const response = await axios.post(`${process.env.BACKEND_URI}/subcribers`, newSubscriber);
        return NextResponse.json(response.data);
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}
