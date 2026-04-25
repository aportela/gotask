import axios from "axios";
import { SERVER_API_BASE_PATH } from "../constants";

const axiosInstance = axios.create({
  baseURL: SERVER_API_BASE_PATH,
});

const bgDownload = async (url: string, fileName: string = "fileName") => {
  const startTime = Date.now();
  const response = await axiosInstance.get(url, {
    responseType: "blob",
  });
  const blob = new Blob([response.data]);
  const tmpLink = document.createElement("a");
  const urlBlob = URL.createObjectURL(blob);
  tmpLink.href = urlBlob;
  tmpLink.download = fileName;
  document.body.appendChild(tmpLink);
  tmpLink.click();
  document.body.removeChild(tmpLink);
  URL.revokeObjectURL(urlBlob);
  const endTime = Date.now();
  return {
    fileName: fileName,
    url: url,
    mimeType: blob.type,
    length: blob.size,
    msTime: endTime - startTime,
  };
};

export { axiosInstance, bgDownload };
