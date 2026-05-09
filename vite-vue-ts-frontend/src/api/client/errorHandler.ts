import axios from "axios";
import type { APIError } from "./errors";

type APIErrorHandler = (error: APIError) => void;
type FatalErrorHandler = (error: unknown) => void;

export const handleAPIError = (
  error: unknown,
  apiErrorCallback?: APIErrorHandler,
  fatalHandler?: FatalErrorHandler,
) => {
  const handleFatal = () => {
    if (fatalHandler) {
      fatalHandler(error);
    } else {
      console.error(error);
    }
  };

  if (axios.isAxiosError(error)) {
    if ((error as APIError).isAPIError) {
      const apiError = error as APIError;
      if (apiErrorCallback) {
        apiErrorCallback(apiError);
      } else {
        handleFatal();
      }
    } else {
      handleFatal();
    }
  } else {
    handleFatal();
  }
};
