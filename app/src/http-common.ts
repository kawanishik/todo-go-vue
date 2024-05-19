import axios from "axios";
import type { AxiosInstance } from "axios";

const http: AxiosInstance = axios.create({
  baseURL: "http://localhost:3000/api",
  headers: {
    "Content-type": "application/json",
  },
});

export default http;