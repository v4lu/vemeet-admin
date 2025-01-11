'use client';
import { useSidebarStore } from '@/app/store/sidebar.store';
import { Icon } from '@iconify/react';

export function Sidebar() {
  const menuItems = [
    { name: 'Dashboard', icon: 'solar:home-2-bold-duotone', active: true },
    { name: 'Analytics', icon: 'solar:chart-2-bold-duotone' },
    { name: 'Users', icon: 'solar:users-group-rounded-bold-duotone' },
    { name: 'Settings', icon: 'solar:settings-bold-duotone' },
  ];
  const { isOpen } = useSidebarStore();

  return (
    <aside
      className={`fixed left-0 top-16 h-[calc(100vh-4rem)] bg-card border-r border-border w-64 transform transition-transform duration-200 ease-in-out ${
        isOpen ? 'translate-x-0' : '-translate-x-full'
      }`}
    >
      <nav className="p-4 space-y-2">
        {menuItems.map((item) => (
          <button
            key={item.name}
            className={`w-full flex items-center gap-2 p-2 rounded-md ${
              item.active
                ? 'bg-primary text-primary-foreground'
                : 'hover:bg-secondary text-foreground'
            }`}
          >
            <Icon icon={item.icon} className="w-5 h-5" />
            {item.name}
          </button>
        ))}
      </nav>
    </aside>
  );
}
