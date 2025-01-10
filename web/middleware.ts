import { NextRequest, NextResponse } from "next/server";

export const config = {
  matcher: "/api/:path*",
};

export function middleware(request: NextRequest) {
  const baseURL = process.env.API_BASE_URL;
  const path = request.nextUrl.pathname.slice("/api".length);
  const proxied = new URL(`${baseURL}${path}${request.nextUrl.search}`);

  return NextResponse.rewrite(proxied, { request });
}
