'use client'

import ListPosts from "@/components/ListPosts";
import Image from "next/image"
import Link from "next/link";
import { useFormik } from 'formik';
import axios from 'axios';
import { useState } from "react";
import * as Yup from 'yup';
import useSWR from 'swr';
import { AiOutlineLoading3Quarters } from "react-icons/ai";

export default function Home() {
  const posts = useSWR(`/api/home_posts`, url => axios.get(url).then(res => res.data))

  if (posts.error) console.log(posts.error)

  const [subcribeLoading, setSubcribeLoading] = useState(false)
  const [subscribeError, setSubscribeError] = useState('')

  const formik = useFormik({
    initialValues: {
      email: '',
      name: '',
    },
    validationSchema: Yup.object({
      name: Yup.string().max(20, 'Must be 20 characters or less').required('Required'),
      email: Yup.string().email('Invalid email address').required('Required'),
    }),
    onSubmit: values => {
      setSubcribeLoading(true)
      axios.post('/api/subcribe', values)
        .then((res) => {
          if (res.data.error) setSubscribeError(`${res.data.error.message}: ${res.data.error.detail}`)
          setSubcribeLoading(false)
        })
        .catch(err => console.log(err))
    },
  });

  return (
    <div>
      <div className="w-full relative h-96">
        <Image
          className="object-cover object-center rounded"
          src="/banner.jpeg"
          alt="Banner"
          fill
        />
      </div>
      <div className="flex w-full mx-auto max-w-container pt-4">
        <div className="flex-3/1 pr-4">
          <ListPosts loading={posts.isLoading} posts={posts.data} />
          {
            posts && posts.data && posts.data.pagination.limit * posts.data.pagination.page < posts.data.pagination.total ?
              <Link href="/blogs" className="underline mt-2 block float-end">Show more &gt;&gt;</Link> : null
          }
        </div>
        <div className="flex-1">
          <h3 className="text-center mb-4 text-3xl font-medium">About me</h3>
          <div className="relative h-56 w-56 mx-auto p-4">
            <Image
              src="/profile.jpg"
              alt="Profile"
              className="rounded-full object-cover object-center"
              fill
            />
          </div>
          <div className="text-justify w-fit mx-auto text-lg mt-4">
            Brief information about me. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut lobortis ultricies accumsan. Sed mattis enim non egestas aliquet. Maecenas iaculis elit nec dolor porta ultricies eget nec ante.
          </div>
        </div>
      </div>
      <div className="bg-gray-800 rounded mt-4 p-4 text-white text-base">
        <h3 className="text-3xl font-bold text-center">Subcribe for receiving newsletter</h3>
        <h5 className="text-xl font-medium text-center mt-2 text-gray-300">Leave the email that I can notify you whenever new blog posted!</h5>
        <form className="md:w-1/2 mx-auto mt-8" onSubmit={formik.handleSubmit}>
          <div className="mb-4">
            <label htmlFor="email" className="mr-2">Email:</label>
            <input
              id="email"
              type="email"
              name="email"
              placeholder="Your email here. Example: abc@xyz.com"
              className="text-sm py-2 px-4 bg-transparent border-b border-gray-300 w-full outline-none"
              onChange={formik.handleChange}
              value={formik.values.email}
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="name" className="mr-2">Name:</label>
            <input
              id="name"
              type="name"
              name="name"
              placeholder="How I can call you? Example: Mr. Example"
              className="text-sm py-2 px-4 bg-transparent border-b border-gray-300 w-full outline-none"
              onChange={formik.handleChange}
              value={formik.values.name}
              required
            />
          </div>
          <div className="text-sm text-red-400">
            {subscribeError && <h3>{subscribeError}</h3>}
          </div>
          <div className="flex justify-end">
            <button className="text-gray-800 rounded font-semibold p-2 bg-gray-200 flex items-center" type="submit">
              Subcribe
              {subcribeLoading && <AiOutlineLoading3Quarters className="animate-spin ml-2 text-lg" />}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
