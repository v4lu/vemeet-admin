import { Header } from '@/components/layout/header';
import { Sidebar } from '@/components/layout/sidebar';
import { SessionProvider } from '@/components/session-provider';
import { authAPI } from '@/lib/api';
import { Admin } from '@/lib/types/types';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import type { ReactNode } from 'react';

async function getSession(): Promise<
  { admin: Admin; accessToken: string } | undefined
> {
  const cookieStore = await cookies();
  const accesToken = cookieStore.get('ACCESS_TOKEN');

  if (!accesToken) {
    redirect('/sign-in');
  }

  try {
    const api = authAPI(accesToken.value);

    const res = await api.get<Admin>('auth').json<Record<'admin', Admin>>();
    return {
      admin: res.admin,
      accessToken: accesToken.value,
    };
  } catch (error) {
    console.error(error);
  }
}
export default async function AppLayout({ children }: { children: ReactNode }) {
  const session = await getSession();
  if (!session) {
    redirect('/sign-in');
  }

  return (
    <SessionProvider user={session.admin} accessToken={session.accessToken}>
      <div className="min-h-screen bg-background">
        <Header />
        <Sidebar />
        {children}
      </div>
    </SessionProvider>
  );
}
