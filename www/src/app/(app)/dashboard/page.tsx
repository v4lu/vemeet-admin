'use client';
import { useSessionStore } from '@/app/store/session.store';
import { Icon } from '@iconify/react';
import { useUsers } from '@/hooks/use-users';
import { useSidebarStore } from '@/app/store/sidebar.store';

export default function Dashboard() {
  const { accessToken } = useSessionStore();
  const { isOpen } = useSidebarStore();

  const { data: userData, isLoading: usersLoading } = useUsers({
    accessToken,
    sort: 'created_at',
    order: 'desc',
  });

  return (
    <main
      className={`pt-16 transition-all duration-200 ${
        isOpen ? 'ml-64' : 'ml-0'
      }`}
    >
      <div className="p-6">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          {usersLoading
            ? Array(4)
                .fill(0)
                .map((_, i) => (
                  <div
                    key={i}
                    className="h-32 bg-card/50 backdrop-blur-xl rounded-xl border border-border/50 p-4 animate-pulse"
                  >
                    <div className="w-12 h-12 rounded-xl bg-muted mb-2"></div>
                    <div className="h-4 bg-muted rounded-lg w-1/2 mb-2"></div>
                    <div className="h-4 bg-muted rounded-lg w-1/3"></div>
                  </div>
                ))
            : [
                {
                  title: 'Total Users',
                  value: '1,234',
                  change: '+12%',
                  icon: 'mdi:users',
                },
                {
                  title: 'Revenue',
                  value: '$12,345',
                  change: '+8%',
                  icon: 'material-symbols:payments',
                },
                {
                  title: 'Active Sessions',
                  value: '432',
                  change: '+5%',
                  icon: 'material-symbols:monitoring',
                },
                {
                  title: 'Conversion Rate',
                  value: '2.3%',
                  change: '+1%',
                  icon: 'material-symbols:trending-up',
                },
              ].map((stat, i) => (
                <div
                  key={i}
                  className="bg-card rounded-lg border border-border p-4 hover:shadow-lg transition-shadow"
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-muted-foreground">{stat.title}</p>
                      <h3 className="text-2xl font-bold text-foreground">
                        {stat.value}
                      </h3>
                      <span className="text-primary">{stat.change}</span>
                    </div>
                    <Icon
                      icon={stat.icon}
                      className="w-12 h-12 text-primary/20"
                    />
                  </div>
                </div>
              ))}
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6"></div>

          <div className="bg-card/50 backdrop-blur-xl rounded-xl border border-border/50 p-6 h-fit">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-lg font-semibold text-foreground">
                Recent Users
              </h2>
              <span className="text-sm text-muted-foreground">
                {userData?.total || 0} total users
              </span>
            </div>

            {usersLoading ? (
              <div className="space-y-4">
                {Array(7)
                  .fill(0)
                  .map((_, i) => (
                    <div
                      key={i}
                      className="flex items-center gap-4 animate-pulse"
                    >
                      <div className="w-12 h-12 rounded-xl bg-muted"></div>
                      <div className="flex-1">
                        <div className="h-4 bg-muted rounded-lg w-1/3 mb-2"></div>
                        <div className="h-4 bg-muted rounded-lg w-1/2"></div>
                      </div>
                    </div>
                  ))}
              </div>
            ) : (
              <div className="space-y-4">
                {userData?.users.slice(0, 7).map((user) => (
                  <div
                    key={user.id}
                    className="group flex items-center gap-4 p-3 rounded-xl hover:bg-secondary/50 transition-colors duration-200"
                  >
                    <div className="w-12 h-12 rounded-xl bg-primary/10 text-primary flex items-center justify-center ring-2 ring-primary/10">
                      {user.profile_image ? (
                        <img
                          src={user.profile_image.url}
                          alt={user.username}
                          className="w-full h-full rounded-xl object-cover"
                        />
                      ) : (
                        <Icon
                          icon="solar:user-bold-duotone"
                          className="w-7 h-7"
                        />
                      )}
                    </div>
                    <div className="flex-1">
                      <div className="flex items-center gap-2">
                        <span className="font-medium text-foreground">
                          {user.name || user.username}
                        </span>
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
                      <div className="flex items-center gap-3 text-sm text-muted-foreground">
                        <span className="flex items-center gap-1">
                          <Icon
                            icon="solar:calendar-bold-duotone"
                            className="w-4 h-4"
                          />
                          {new Date(user.created_at).toLocaleDateString()}
                        </span>
                        {user.country_name && (
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
                        )}
                      </div>
                    </div>
                    <button className="opacity-0 group-hover:opacity-100 p-2 rounded-lg hover:bg-secondary transition-all duration-200">
                      <Icon
                        icon="solar:alt-arrow-right-bold-duotone"
                        className="w-5 h-5 text-foreground"
                      />
                    </button>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </main>
  );
}
