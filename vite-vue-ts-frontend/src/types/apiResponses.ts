import { type AxiosResponse } from "axios";

import { type ValidAuthTypes } from "./common";

interface DefaultAxiosResponse<T = unknown> {
  data: AxiosResponse<T>;
}

interface SignInSucessResponse extends Omit<DefaultAxiosResponse, "data"> {
  data: {
    accessToken: {
      token: string;
      expiresAtTimestamp: number;
    };
    refreshToken: {
      token: string;
      expiresAtTimestamp: number;
    };
    tokenType: ValidAuthTypes;
  };
}

interface GetNewAccessTokenResponse extends Omit<DefaultAxiosResponse, "data"> {
  data: {
    accessToken: {
      token: string;
      expiresAtTimestamp: number;
    };
    tokenType: ValidAuthTypes;
  };
}

export {
  type DefaultAxiosResponse,
  type SignInSucessResponse,
  type GetNewAccessTokenResponse,
};
