/*
import { authService } from "./auth";
import { useSessionStore } from "../../stores/session";

export const TokenManager = {
  async refreshAccessToken(): Promise<boolean> {
    try {
      const sessionStore = useSessionStore();
      const response = await authService.renewAccessToken();
      sessionStore.setAccessToken(
        response.accessToken.token,
        response.accessToken.expiresAtTimestamp,
      );
      return true;
    } catch (err) {
      console.error("Failed to refresh token", err);
      return false;
    }
  },
};
*/

import { authService } from "./auth";
import { useSessionStore } from "../../../stores/session";
import type { RenewAccessTokenResponse } from "../types/dto";
import { User } from "../../users/models/user";

export const TokenManager = {
  async refreshAccessToken(
    sessionStore: ReturnType<typeof useSessionStore>,
  ): Promise<boolean> {
    try {
      const response: RenewAccessTokenResponse =
        await authService.renewAccessToken();
      sessionStore.setAccessToken(
        response.accessToken.token,
        response.accessToken.expiresAtTimestamp,
      );
      sessionStore.setUser(new User(response.user));
      return true;
    } catch (err) {
      console.error("Failed to refresh token", err);
      return false;
    }
  },
};
