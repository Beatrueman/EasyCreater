// src/utils/axios.ts
import axios from 'axios';

// 创建 Axios 实例
const axiosInstance = axios.create({
  //baseURL: import.meta.env.VITE_API_SERVER || 'http://localhost:8888/api',
  baseURL: '/api',
  timeout: 50000,  // 请求超时时间
});

// 请求拦截器）
axiosInstance.interceptors.request.use(
  (config) => {
    // 获取本地存储的 token
    const token = localStorage.getItem('jwt-token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器（处理后端响应）
axiosInstance.interceptors.response.use(
  (response) => response.data,  // 直接返回数据
  (error) => {
    return Promise.reject(error);
  }
);

export default axiosInstance;
