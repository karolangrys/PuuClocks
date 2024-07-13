import axios from 'axios';

export const BackendApi = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});
