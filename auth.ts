import NextAuth, { NextAuthConfig } from 'next-auth';
import Google from 'next-auth/providers/google';

export const config: NextAuthConfig = {
  providers: [Google],
  pages: {
    signIn: '/auth/login',
  },
};

export const { signIn, handlers, auth } = NextAuth(config);
