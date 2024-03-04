'use client'

import * as Yup from 'yup';
import { useFormik } from 'formik';
import axios from 'axios';
import { useState, useContext, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { UserContext } from '@/contexts/UserContext';

export default function Home() {
  const [error, setError] = useState('')
  const { user, login } = useContext(UserContext); 
  const router = useRouter();
  useEffect(()=>{
    if (user) router.push('/dashboard')
  }, [user, router])
  const formik = useFormik({
      initialValues: {
        username: '',
        password: '',
      },
      validationSchema: Yup.object({
        username: Yup.string().required('Required'),
        password: Yup.string().required('Required'),
      }),
      onSubmit: values => {
        axios.post('/api/verify', values)
        .then(res => {
          if (res.data.error) setError(`${res.data.error.message}: ${res.data.error.detail}`)
          else {
            login(res.data.data.token)
            router.push("/dashboard")
          }
        })
        .catch(err => console.log(err))
      },
  });
  return (
      <div className='flex items-center justify-center w-full h-screen'>
          <form className='w-96 bg-gray-100 px-16 py-8 rounded' onSubmit={formik.handleSubmit}>
              <h1 className='text-xl font-semibold uppercase mb-4'>Welcome back, admin!</h1>
              <div className="mb-4">
                  <label htmlFor="username" className="mr-2">Username:</label>
                  <input
                  id="username"
                  type="text"
                  name="username"
                  placeholder="Username"
                  className="text-sm py-2 px-4 bg-transparent border-b border-gray-300 w-full outline-none"
                  onChange={formik.handleChange}
                  value={formik.values.username}
                  required
                  />
              </div>
              <div className="mb-4">
                  <label htmlFor="username" className="mr-2">Password:</label>
                  <input
                  id="password"
                  type="password"
                  name="password"
                  placeholder="Password"
                  className="text-sm py-2 px-4 bg-transparent border-b border-gray-300 w-full outline-none"
                  onChange={formik.handleChange}
                  value={formik.values.password}
                  required
                  />
              </div>
              <div className="text-sm text-red-400 mb-4">
                {error && <h3>{error}</h3>}
              </div>
              <div className='flex justify-end'>
                <button type='submit' className='bg-gray-400 text-white px-4 py-2 rounded'>Log in</button>
              </div>
          </form>
      </div>
  )
}