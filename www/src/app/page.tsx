'use client';
import React from 'react';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { signInForm, SignInForm } from '@/lib/validators';
import { signIn } from '@/lib/actions';

export default function SignIn() {
  const { register, handleSubmit } = useForm<SignInForm>({
    resolver: zodResolver(signInForm),
  });

  function onSubmit(data: SignInForm) {
    signIn(data);
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-background">
      <div className="w-full max-w-xl p-8 space-y-6 bg-card rounded-lg shadow-2xl border border-border">
        <div className="space-y-2 text-center">
          <h2 className="text-4xl font-bold text-primary">Admin Area</h2>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <label
              htmlFor="email"
              className="text-sm font-medium text-foreground"
            >
              Email
            </label>
            <input
              id="email"
              placeholder="Enter your email"
              className="w-full px-4 py-2 border border-input rounded-md bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:border-transparent transition-colors"
              {...register('email')}
            />
          </div>

          <div className="space-y-2">
            <label
              htmlFor="password"
              className="text-sm font-medium text-foreground"
            >
              Password
            </label>
            <input
              id="password"
              type="password"
              placeholder="Enter your password"
              className="w-full px-4 py-2 border border-input rounded-md bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:border-transparent transition-colors"
              {...register('password')}
            />
          </div>

          <button
            className="w-full px-4 py-2 text-sm font-medium text-primary-foreground bg-primary rounded-md hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 transition-colors"
            type="submit"
          >
            Sign In
          </button>
        </form>

        <div className="text-center text-sm text-muted-foreground">
          <p className="hover:text-primary transition-colors">
            Forgot your password? Contact support
          </p>
        </div>
      </div>
    </div>
  );
}
