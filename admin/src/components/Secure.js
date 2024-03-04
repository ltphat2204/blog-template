'use client'

import { useRouter } from 'next/navigation';
import { useContext, useEffect } from 'react';
import { UserContext } from '../contexts/UserContext';

export default function RedirectToLoginIfNotAuthenticated({ children }){
    const router = useRouter();
    const { user } = useContext(UserContext);
  
    useEffect(() => {
      if (!user) {
        router.push('/');
      }
    }, [user, router]);
  
    return user ? children : null;
  }
