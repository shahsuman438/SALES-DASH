import axios, { AxiosInstance } from 'axios';
import config from '../config/config';

const axiosInstance: AxiosInstance = axios.create({
  baseURL: config.apiEndpoint,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request Interceptor
axiosInstance.interceptors.request.use(
  (config) => {
    // add token if required
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// Response Interceptor
axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // TODO handle refresh token if there is token expire error
    // TODO handle unauthorization error if the user is unauthenticated > logout > clear token > redirect to login

    const { status } = error.response;
    if (status === 401 || status === 403) {
      handleUnauthorized();
      return;
    }
    return Promise.reject(error);
  },
);

const handleUnauthorized = () => {
  setTimeout(() => {
    window.location.reload();
  }, 1000);
};
export default axiosInstance;
