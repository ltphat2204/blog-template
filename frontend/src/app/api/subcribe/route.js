import axios from 'axios';
import { NextResponse } from 'next/server';

export async function POST(req) {
    const newSubcriber = await req.json()
    const response = await axios.post(`${process.env.BACKEND_URI}/subcribe`, newSubcriber);
    return NextResponse.json(response)
}