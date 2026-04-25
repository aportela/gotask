import { axiosInstance } from "./axios";

const api = {
  user: {
    search: () => axiosInstance.get("/api/users"),
  },
  project: {
    search: () => axiosInstance.get("/api/projects"),
  },
};

export { api };
