import Fetcher from "@/common/Fetcher";
import { NextResponse } from 'next/server';

export async function GET(req) {
    const searchParams = req.nextUrl.searchParams;
    const offset = searchParams.get('offset')
    const max = searchParams.get('max')
    const res = await Fetcher(`${process.env.BACKEND_URI}/posts?offset=${offset}&max=${max}`);
    return NextResponse.json(res)
    
}