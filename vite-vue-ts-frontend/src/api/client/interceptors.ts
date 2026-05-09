import type {
  AxiosError,
  InternalAxiosRequestConfig,
  AxiosHeaders,
} from "axios";
import { axiosInstance } from "./client";
import { useSessionStore } from "../../stores/session";
import type { AxiosAPIError } from "./errors";
import { enrichAxiosError } from "./errors";

axiosInstance.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  const sessionStore = useSessionStore();

  config.headers = config.headers ?? {};

  if (sessionStore.hasAccessToken) {
    config.headers.Authorization = `Bearer ${sessionStore.accessToken}`;
  }

  return config;
});

const getContentType = (
  headers: AxiosHeaders | Record<string, unknown> | undefined,
): string | undefined => {
  if (!headers) return undefined;
  if ("get" in headers && typeof headers.get === "function") {
    const ct = headers.get("content-type");
    return typeof ct === "string" ? ct : undefined;
  } else {
    const raw = (headers as Record<string, unknown>)["content-type"];
    if (Array.isArray(raw)) {
      return raw.find((v) => typeof v === "string");
    } else if (typeof raw === "string") {
      return raw;
    } else {
      return undefined;
    }
  }
};

axiosInstance.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    const sessionStore = useSessionStore();

    if (error.response?.status === 401) {
      if (sessionStore.hasAccessToken) {
        sessionStore.removeAccessToken();
      }
    }

    const apiError = enrichAxiosError(error) as AxiosAPIError;

    const contentType = getContentType(error.response?.headers);
    apiError.isAPIError = !!(
      contentType?.toLowerCase().includes("application/json") &&
      (error.response?.data as any)?.APIError === true
    );

    return Promise.reject(apiError);
  },
);
