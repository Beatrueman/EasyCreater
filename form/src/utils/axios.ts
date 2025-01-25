// src/utils/axios.ts
import axios from 'axios';

// 创建 Axios 实例
const axiosInstance = axios.create({
  baseURL: 'http://localhost:8888',  // 后端 API 地址
  timeout: 5000,  // 请求超时时间
});

// 请求拦截器（如果需要加上 token 等信息）
axiosInstance.interceptors.request.use(
  (config) => {
    // 获取本地存储的 token
    const token = localStorage.getItem('token');
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
