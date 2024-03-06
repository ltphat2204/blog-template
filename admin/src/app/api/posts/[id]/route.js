import axios from 'axios';
import { NextResponse } from 'next/server';

export async function GET(request, {params}) {
    try {
        const url = `${process.env.BACKEND_URI}/posts/${params.id}`;

        const res = await axios.get(url);
        const data = res.data;

        return NextResponse.json(data);
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}

export async function PATCH(request, {params}) {
    try {
        const post = await request.json();
        const url = `${process.env.BACKEND_URI}/posts/${params.id}`;

        const res = await axios.patch(url, post, {
            headers: {
                Authorization: request.headers.get('authorization')
            }
        });
        const data = res.data;

        return NextResponse.json(data);
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}

export async function DELETE(request, {params}) {
    try {
        const url = `${process.env.BACKEND_URI}/posts/${params.id}`;

        const res = await axios.delete(url, {
            headers: {
                Authorization: request.headers.get('authorization')
            }
        });
        const data = res.data;

        return NextResponse.json(data);
    } catch (error) {
        console.error('Error:', error.response.data);
        return NextResponse.json(error.response.data);
    }
}
