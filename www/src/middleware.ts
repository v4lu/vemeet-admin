import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
import { AuthResponse } from './lib/types/types';
import { api } from './lib/api';

export async function middleware(request: NextRequest) {
  const access_token = request.cookies.get('ACCESS_TOKEN');
  const refresh_token = request.cookies.get('REFRESH_TOKEN');

  if (!refresh_token) {
    return NextResponse.redirect('/sign-in');
  }

  if (!access_token) {
    try {
      const res = await api
        .post('auth/refresh-token', {
          headers: {
            'Refresh-Token-X': refresh_token.value,
          },
        })
        .json<AuthResponse>();

      request.cookies.delete('REFRESH_TOKEN');

      const response = NextResponse.next();
      const accessTokenMaxAge = Math.floor(
        (new Date(res.tokens.access_token_expiry).getTime() - Date.now()) / 1000
      );
      const refreshTokenMaxAge = Math.floor(
        (new Date(res.tokens.refresh_token_expiry).getTime() - Date.now()) /
          1000
      );

      response.cookies.set('ACCESS_TOKEN', res.tokens.access_token, {
        maxAge: accessTokenMaxAge,
        path: '/',
        secure: true,
        httpOnly: true,
        sameSite: 'strict',
      });
      response.cookies.set('REFRESH_TOKEN', res.tokens.refresh_token, {
        maxAge: refreshTokenMaxAge,
        path: '/',
        secure: true,
        httpOnly: true,
        sameSite: 'strict',
      });
      return response;
    } catch (error) {
      console.error(error);
    }
  }

  return NextResponse.next();
}
export const config = {
  matcher: ['/dashboard'],
};
