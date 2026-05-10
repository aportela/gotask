import { axiosInstance } from "../client";
import type {
  SignInResponse,
  RenewAccessTokenResponse,
} from "../types/dto/auth";

export const authService = {
  async signIn(email: string, password: string): Promise<SignInResponse> {
    const params = { email, password };
    const { data } = await axiosInstance.post<SignInResponse>(
      "/auth/signin",
      params,
    );
    return data;
  },
  async signOut(): Promise<void> {
    await axiosInstance.post("/auth/signout");
  },
  async renewAccessToken(): Promise<RenewAccessTokenResponse> {
    const { data } = await axiosInstance.post<RenewAccessTokenResponse>(
      "/auth/renew-access-token",
    );
    return data;
  },
};
