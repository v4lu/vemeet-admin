import { Admin } from '@/lib/types/types';
import { create } from 'zustand';

type SessionStore = {
  admin: Admin | null;
  setUser: (admin: Admin) => void;
  setAccessToken: (accessToken: string) => void;
  deleteUser: () => void;
  isLoading: boolean;
  accessToken: string;
};

export const useSessionStore = create<SessionStore>((set) => ({
  accessToken: '',
  setAccessToken: (accessToken: string) => {
    set({ accessToken });
  },
  admin: null,
  isLoading: true,
  setUser: (admin) => {
    set({ admin, isLoading: false });
  },
  deleteUser: () => set({ admin: undefined }),
}));
