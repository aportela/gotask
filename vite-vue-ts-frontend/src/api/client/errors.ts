import type { AxiosError } from "axios";

export interface APIErrorDetails {
  method: string;
  url: string;

  http: {
    code: number;
    status: string;
    contentType?: string;
  };

  request: {
    params?: unknown;
    data?: unknown;
  };

  response: unknown;
}

export type AxiosAPIError = AxiosError & {
  isAPIError: boolean;
  customAPIErrorDetails: APIErrorDetails;
};

export const enrichAxiosError = (error: AxiosError): AxiosAPIError => {
  const apiError = error as AxiosAPIError;

  const rawContentType = error.response?.headers?.["content-type"];

  const contentType = Array.isArray(rawContentType)
    ? rawContentType[0]
    : typeof rawContentType === "string"
      ? rawContentType
      : undefined;

  apiError.customAPIErrorDetails = {
    method: error.config?.method ?? "unknown",
    url: error.config?.url ?? error.request?.responseURL ?? "unknown",

    http: {
      code: error.response?.status ?? 0,
      status: error.response?.statusText ?? "Unknown error",
      contentType,
    },

    request: {
      params: error.config?.params,
      data: error.config?.data,
    },

    response: error.response?.data,
  };

  return apiError;
};
