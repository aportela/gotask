import axios from "axios";
import { SERVER_API_BASE_PATH } from "../../constants";

export const axiosInstance = axios.create({
  baseURL: SERVER_API_BASE_PATH,
  withCredentials: true,
});
