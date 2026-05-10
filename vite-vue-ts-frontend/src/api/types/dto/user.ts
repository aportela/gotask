/*
export type UserRequest = {
  id: string;
  name: string;
  email: string;
  isSuperUser: boolean;
  createdAt: number;
  updatedAt: number | null;
  deletedAt: number | null;
  avatarURL: string;
};
*/

export type AddRequestInterface = {
  user: {
    name: string;
    email: string;
    isSuperUser: boolean;
  };
};

export type UpdateRequestInterface = {
  user: {
    id: string;
    name: string;
    email: string;
    isSuperUser: boolean;
  };
};

export type SearchRequestInterface = {
  name?: string;
  email?: string;
  isSuperUser?: boolean;
  page?: number;
  limit?: number;
};

export type UserResponse = {
  id: string;
  name: string;
  email: string;
  isSuperUser: boolean;
  createdAt: number;
  updatedAt: number | null;
  deletedAt: number | null;
  avatarURL: string | null;
};

export type AddResponseInterface = {
  user: UserResponse;
};

export type UpdateResponseInterface = {
  user: UserResponse;
};

export type GetResponseInterface = {
  user: UserResponse;
};

export type SearchResponseInterface = {
  users: UserResponse[];
};
