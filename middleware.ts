import { auth } from '@/auth';
import { NextResponse } from 'next/server';

const LOGIN_ROUTE = '/auth/login';
const HOME_ROUTE = '/';

export default auth((req) => {
  const isAuthenticated = !!req.auth;
  const { nextUrl } = req;

  if (!isAuthenticated && !nextUrl.pathname.startsWith(LOGIN_ROUTE))
    return NextResponse.redirect(new URL(LOGIN_ROUTE, nextUrl.origin));

  if (isAuthenticated && nextUrl.pathname.startsWith(LOGIN_ROUTE))
    return NextResponse.redirect(new URL(HOME_ROUTE, nextUrl.origin));
});

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
