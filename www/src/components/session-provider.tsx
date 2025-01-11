'use client';

import { useSessionStore } from '@/app/store/session.store';
import { Admin } from '@/lib/types/types';
import { ReactNode, useEffect } from 'react';

export function SessionProvider({
  user,
  children,
  accessToken,
}: {
  user: Admin;
  accessToken: string;
  children: ReactNode;
}) {
  const { setUser, setAccessToken } = useSessionStore();

  useEffect(() => {
    setUser(user);
    setAccessToken(accessToken);
  }, [user, setUser, accessToken, setAccessToken]);

  return <> {children}</>;
}
