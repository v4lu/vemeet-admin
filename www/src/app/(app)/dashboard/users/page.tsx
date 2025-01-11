'use client';
import { useState } from 'react';
import { Icon } from '@iconify/react';
import { useSessionStore } from '@/app/store/session.store';
import { User } from '@/lib/types/user.type';
import { useUsers } from '@/hooks/use-users';

export default function UsersPage() {
  const { accessToken } = useSessionStore();
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState('');
  const [sort, setSort] = useState('created_at');
  const [order, setOrder] = useState<'asc' | 'desc'>('desc');

  const { data: userData, isLoading } = useUsers({
    accessToken: accessToken,
    page,
    sort,
    order,
    search,
  });

  const sortOptions = [
    { label: 'Date Created', value: 'created_at' },
    { label: 'Username', value: 'username' },
    { label: 'Name', value: 'name' },
  ];

  return (
    <div className="p-6  mx-auto">
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-foreground mb-4">
          Users Management
        </h1>
        <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
          <div className="relative w-full sm:w-96">
            <Icon
              icon="solar:magnifier-bold-duotone"
              className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground"
            />
            <input
              type="text"
              placeholder="Search users..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-full pl-10 pr-4 py-2 rounded-xl bg-card/50 border border-border/50 focus:outline-none focus:ring-2 focus:ring-primary/20"
            />
          </div>

          <div className="flex gap-3 w-full sm:w-auto">
            <select
              value={sort}
              onChange={(e) => setSort(e.target.value)}
              className="px-4 py-2 rounded-xl bg-card/50 border border-border/50 focus:outline-none focus:ring-2 focus:ring-primary/20"
            >
              {sortOptions.map((option) => (
                <option key={option.value} value={option.value}>
                  {option.label}
                </option>
              ))}
            </select>
            <button
              onClick={() => setOrder(order === 'asc' ? 'desc' : 'asc')}
              className="p-2 rounded-xl bg-card/50 border border-border/50 hover:bg-secondary"
            >
              <Icon
                icon={
                  order === 'asc'
                    ? 'solar:sort-from-bottom-to-top-bold-duotone'
                    : 'solar:sort-from-top-to-bottom-bold-duotone'
                }
                className="w-5 h-5 text-foreground"
              />
            </button>
          </div>
        </div>
      </div>

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {Array(8)
            .fill(0)
            .map((_, i) => (
              <div
                key={i}
                className="bg-card/50 backdrop-blur-xl rounded-xl border border-border/50 p-6 animate-pulse"
              >
                <div className="flex items-center gap-4 mb-4">
                  <div className="w-16 h-16 rounded-xl bg-muted"></div>
                  <div className="flex-1">
                    <div className="h-4 bg-muted rounded-lg w-2/3 mb-2"></div>
                    <div className="h-4 bg-muted rounded-lg w-1/2"></div>
                  </div>
                </div>
                <div className="space-y-2">
                  <div className="h-4 bg-muted rounded-lg w-full"></div>
                  <div className="h-4 bg-muted rounded-lg w-3/4"></div>
                </div>
              </div>
            ))}
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {userData?.users.map((user: User) => (
              <div
                key={user.id}
                className="group bg-card/50 backdrop-blur-xl rounded-xl border border-border/50 p-6 hover:shadow-xl hover:shadow-primary/5 transition-all duration-300"
              >
                <div className="flex items-center gap-4 mb-4">
                  <div className="w-16 h-16 rounded-xl bg-primary/10 ring-2 ring-primary/20 flex items-center justify-center overflow-hidden">
                    {user.profile_image ? (
                      <img
                        src={user.profile_image.url}
                        alt={user.username}
                        className="w-full h-full object-cover"
                      />
                    ) : (
                      <Icon
                        icon="solar:user-bold-duotone"
                        className="w-8 h-8 text-primary"
                      />
                    )}
                  </div>
                  <div>
                    <div className="flex items-center gap-2">
                      <h3 className="font-semibold text-foreground">
                        {user.name || user.username}
                      </h3>
                      {user.verified && (
                        <Icon
                          icon="solar:verified-check-bold-duotone"
                          className="w-4 h-4 text-primary"
                        />
                      )}
                      {user.is_private && (
                        <Icon
                          icon="solar:lock-bold-duotone"
                          className="w-4 h-4 text-muted-foreground"
                        />
                      )}
                    </div>
                    <p className="text-sm text-muted-foreground">
                      @{user.username}
                    </p>
                  </div>
                </div>

                <div className="space-y-2">
                  {user.country_name && (
                    <div className="flex items-center gap-2 text-sm text-muted-foreground">
                      <Icon
                        icon="solar:globe-bold-duotone"
                        className="w-4 h-4"
                      />
                      <span className="flex items-center gap-1">
                        {user.country_flag && (
                          <img
                            src={user.country_flag}
                            alt={user.country_name}
                            className="w-4 h-4"
                          />
                        )}
                        {user.country_name}
                      </span>
                    </div>
                  )}
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Icon
                      icon="solar:calendar-bold-duotone"
                      className="w-4 h-4"
                    />
                    Joined {new Date(user.created_at).toLocaleDateString()}
                  </div>
                </div>

                <div className="flex gap-2 mt-4 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button className="flex-1 p-2 rounded-lg bg-primary/10 hover:bg-primary/20 text-primary text-sm font-medium">
                    Edit
                  </button>
                  <button className="p-2 rounded-lg hover:bg-secondary">
                    <Icon
                      icon="solar:menu-dots-bold-duotone"
                      className="w-5 h-5 text-foreground"
                    />
                  </button>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-6 flex items-center justify-between">
            <p className="text-sm text-muted-foreground">
              Showing {(page - 1) * 10 + 1} to{' '}
              {Math.min(page * 10, userData?.total || 0)} of {userData?.total}{' '}
              users
            </p>
            <div className="flex gap-2">
              <button
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                disabled={page === 1}
                className="p-2 rounded-xl bg-card/50 border border-border/50 hover:bg-secondary disabled:opacity-50"
              >
                <Icon
                  icon="solar:arrow-left-bold-duotone"
                  className="w-5 h-5 text-foreground"
                />
              </button>
              <button
                onClick={() => setPage((p) => p + 1)}
                disabled={!userData?.has_more}
                className="p-2 rounded-xl bg-card/50 border border-border/50 hover:bg-secondary disabled:opacity-50"
              >
                <Icon
                  icon="solar:arrow-right-bold-duotone"
                  className="w-5 h-5 text-foreground"
                />
              </button>
            </div>
          </div>
        </>
      )}
    </div>
  );
}
