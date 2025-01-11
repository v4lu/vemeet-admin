'use server';

import { AuthResponse } from './types/types';
import { SignInForm } from './validators';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import { api } from './api';

export async function signIn(data: SignInForm) {
  const res = await api
    .post('auth/login', {
      json: data,
    })
    .json<AuthResponse>();

  const accessTokenMaxAge = Math.floor(
    (new Date(res.tokens.access_token_expiry).getTime() - Date.now()) / 1000
  );
  const refreshTokenMaxAge = Math.floor(
    (new Date(res.tokens.refresh_token_expiry).getTime() - Date.now()) / 1000
  );

  const cookieStore = await cookies();
  cookieStore.set('ACCESS_TOKEN', res.tokens.access_token, {
    maxAge: accessTokenMaxAge,
    path: '/',
    secure: true,
    httpOnly: true,
    sameSite: 'strict',
  });
  cookieStore.set('REFRESH_TOKEN', res.tokens.refresh_token, {
    maxAge: refreshTokenMaxAge,
    path: '/',
    secure: true,
    httpOnly: true,
    sameSite: 'strict',
  });

  redirect('/dashboard');
}
