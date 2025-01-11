import { authAPI } from '@/lib/api';
import { UserPagination } from '@/lib/types/user.type';
import { useQuery } from '@tanstack/react-query';

async function fetchUsers(
  accessToken: string,
  page = 1,
  sort = 'id',
  order = 'asc',
  search = ''
): Promise<UserPagination | undefined> {
  const api = authAPI(accessToken);
  console.log(accessToken);
  const res = await api
    .get(
      `users?page=${page}&pageSize=10&sort=${sort}&order=${order}&search=${search}`
    )
    .json<UserPagination>();

  return res;
}

export function useUsers({
  accessToken,
  page = 1,
  sort = 'id',
  order = 'asc',
  search = '',
}: {
  accessToken: string;
  page?: number;
  sort?: string;
  order?: 'asc' | 'desc';
  search?: string;
}) {
  return useQuery({
    queryKey: ['users', { page, sort, order, search }],
    queryFn: () => fetchUsers(accessToken, page, sort, order, search),
  });
}
