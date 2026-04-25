import { axiosInstance } from "./axios";

const api = {
  project: {
    search: () => axiosInstance.get("/api/projects"),
  },
};

export { api };
