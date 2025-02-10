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

export const askAI = async (resumeData: string) => {
  try {
    const response = await axiosInstance.post("/user/ask", { resume_data: resumeData }, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log("AI 优化建议:", response);
    return response;
  } catch (error) {
    console.error("AI 请求失败:", error);
    throw error;
  }
};

export const askAIBase = async () => {
  try {
    const response = await axiosInstance.get("/user/ask_base", {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log("AI 优化建议:", response);
    return response;
  } catch (error) {
    console.error("AI 请求失败:", error);
    throw error;
  }
};

export const changeUserPassword = async (form: { Password: string; newPassword: string }) => {
  try {
    const response = await axiosInstance.post("/user/change", form, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("修改密码失败:", error);
    throw error; 
  }
};


export const getUserInfo = async () => {
  try {
    const data = await axiosInstance.get("/user/info", {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log("获取用户信息成功!");
    return data;
  } catch (error) {
    console.error("获取用户信息失败:", error);
    throw error;
  }
};
