import axios from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request) {
    try {
        const {username, password} = await request.json();
        const response = await axios.get(`${process.env.BACKEND_URI}/users/token`, {
            auth: {
                username: username,
                password: password
            }
        });

        const data = response.data;
    
        return NextResponse.json({ data: data.data });
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}