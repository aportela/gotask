import type { UserResponse } from "./user";

export type Token = {
  token: string;
  expiresAtTimestamp: number;
};

type ValidAuthTypes = "Bearer";

export type SignInResponse = {
  accessToken: Token;
  refreshToken: Token;
  tokenType: ValidAuthTypes;
  user: UserResponse;
};

export type RenewAccessTokenResponse = {
  accessToken: Token;
  tokenType: ValidAuthTypes;
  user: UserResponse;
};
