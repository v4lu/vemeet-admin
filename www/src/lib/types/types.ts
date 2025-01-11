export type Admin = {
  id: number;
  name: string;
  email: string;
  verified: boolean;
};

export type Tokens = {
  access_token: string;
  refresh_token: string;
  access_token_expiry: string;
  refresh_token_expiry: string;
};

export type AuthResponse = {
  admin: Admin;
  tokens: Tokens;
};
