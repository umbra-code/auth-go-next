import { NextResponse } from 'next/server'

export function middleware(request) {
  const token = request.cookies.get('access-token')

  if (!token && request.nextUrl.pathname !== '/auth/login' && request.nextUrl.pathname !== '/auth/register') {
    return NextResponse.redirect(new URL('/auth/login', request.url))
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
}