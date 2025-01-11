export type Image = {
  id: number;
  user_id: number;
  url: string;
  created_at: string;
};

export type User = {
  id: number;
  username: string;
  birthday: string;
  aws_cognito_id: string;
  created_at: string;
  verified: boolean;
  is_private: boolean;
  inbox_locked: boolean;
  swiper_mode: boolean;
  name?: string;
  gender?: string;
  country_name?: string;
  country_flag?: string;
  country_iso_code?: string;
  country_lat?: number;
  country_lng?: number;
  city_name?: string;
  city_lat?: number;
  city_lng?: number;
  bio?: string;
  profile_image_id?: number;
  profile_image?: Image;
};

export type UserPagination = {
  users: User[];
  total: number;
  has_more: boolean;
  total_pages: number;
  page: number;
  sort: string;
  order: string;
};
