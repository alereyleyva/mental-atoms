import { signIn } from '@/auth';

export default function LoginPage() {
  const handleSignIn = async () => {
    'use server';

    await signIn('google');
  };

  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      <form action={handleSignIn}>
        <button type={'submit'}>Sign in with Google</button>
      </form>
    </main>
  );
}
