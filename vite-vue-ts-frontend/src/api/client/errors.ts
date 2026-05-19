import type { AxiosError } from "axios";

export type APIErrorResponse = {
  APIError?: boolean;
  code?: string;
  message?: string;
  debug?: string;
  details?: any;
};

export type APIError = AxiosError & {
  isAPIError: boolean;
  code: string;
  message: string;
  debug: string;
  details?: any;
};
