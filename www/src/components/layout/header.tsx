'use client';
import { useSessionStore } from '@/app/store/session.store';
import { Icon } from '@iconify/react';
import { useSidebarStore } from '@/app/store/sidebar.store';

export function Header() {
  const { isLoading, admin } = useSessionStore();
  const { isOpen, toggle } = useSidebarStore();

  return (
    <header className="bg-card border-b border-border h-16 fixed w-full top-0 z-50">
      <div className="flex items-center justify-between h-full px-4">
        <div className="flex items-center gap-4">
          <button
            onClick={() => toggle()}
            className="p-2 rounded-md hover:bg-secondary"
          >
            <Icon
              icon={isOpen ? 'mdi:menu-open' : 'mdi:menu'}
              className="w-6 h-6 text-foreground"
            />
          </button>
          <h1 className="text-xl font-bold text-foreground">Dashboard</h1>
        </div>

        <div className="flex items-center gap-4">
          <button className="p-2 rounded-full hover:bg-secondary relative">
            <Icon icon="mdi:bell" className="w-6 h-6 text-foreground" />
            <span className="absolute top-1 right-1 w-2 h-2 bg-primary rounded-full"></span>
          </button>
          <div className="flex items-center gap-2">
            {isLoading ? (
              <div className="w-8 h-8 rounded-full bg-muted animate-pulse"></div>
            ) : (
              <div className="w-8 h-8 rounded-full bg-primary text-primary-foreground flex items-center justify-center">
                {admin?.name?.charAt(0)}
              </div>
            )}
            <span className="text-foreground">
              {isLoading ? 'Loading...' : admin?.name}
            </span>
          </div>
        </div>
      </div>
    </header>
  );
}
