import ky, { KyInstance } from 'ky';
import { CLIENT_URL } from './constants';

function getBaseUrl(): string {
  return CLIENT_URL;
}

export const api = ky.create({
  prefixUrl: getBaseUrl(),
});

export function authAPI(authToken: string): KyInstance {
  return ky.create({
    prefixUrl: getBaseUrl(),
    headers: {
      Authorization: `Bearer ${authToken}`,
    },
    retry: {
      limit: 2,
      methods: ['get', 'post', 'put', 'delete', 'patch'],
      statusCodes: [500],
    },
  });
}
