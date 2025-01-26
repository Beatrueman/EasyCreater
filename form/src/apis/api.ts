import axiosInstance from '../utils/axios';

interface RegisterData {
  username: string;
  email: string;
  password: string;
  phone: string;
}

interface LoginData {
  username: string;
  password: string;
}

// 响应数据类型
interface LoginResponse {
  msg: string;
  status: number;
}

export const Register = async (data: RegisterData) => {
  try {
    const response = await axiosInstance.post('/register', data, {
        headers: {
          'Content-Type': 'application/json',
        },  
    });
    console.log('注册成功:', response);
    return response;
  } catch (error) {
    console.error('注册失败:', error);
    throw error;
  }
};

// // 你还可以添加登录等其他 API 请求
export const Login = async (data: LoginData) => {
  try {
    const response = await axiosInstance.post('/login', data, {
        headers: {
          'Content-Type': 'application/json',
        },
    });
    console.log('登录成功:', response);
    return response;
  } catch (error) {
    console.error('登录失败:', error);
    throw error;
  }
};
