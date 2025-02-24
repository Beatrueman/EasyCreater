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

export const askAIBase = async (resumeData: string) => {
  try {
    const response = await axiosInstance.post("/user/ask_base", { resume_data: resumeData }, {
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

export const saveResume = async (resumeData: any) => {
  try {
    const response = await axiosInstance.post(`/user/resume/save`, {
      resume_data: JSON.stringify(resumeData), // 转换为 JSON 字符串
    });
    return response.data;
  } catch (error) {
    console.error("保存简历失败:", error);
    throw error;
  }
};

// 获取已保存的简历列表
export const fetchResumeList = async () => {
  try {
    const response = await axiosInstance.get(`/user/resume/list`);
    return response.data; // 返回简历数组
  } catch (error) {
    console.error("获取简历列表失败:", error);
    throw error;
  }
};

export const fetchResume = async (resumeId: number) => {
  try {
    const response = await axiosInstance.get(`/user/resume/${resumeId}`);
    return response.data; 
  } catch (error) {
    console.error("获取简历列表失败:", error);
    throw error;
  }
};

export const deleteResume = async (resumeId: number) => {
  try {
    const response = await axiosInstance.delete(`/user/resume/delete/${resumeId}`);
    return response;
  } catch (error) {
    console.error("获取简历列表失败:", error);
    throw error;
  }
};

export const uploadAvatar = async (ImgBase64String: Base64URLString) => {
  try {
    const response = await axiosInstance.post(`/user/avatar/upload`, {
      "avatar": ImgBase64String,
    });
    return response;
  } catch(error) {
    console.error("上传头像失败:", error);
    throw error;
  }
};

export  const loadAvatar = async () => {
  try {
    const response = await axiosInstance.get(`/user/avatar/load`);
    return response;
  } catch(error) {
    console.error("加载头像失败:", error);
    throw error;
  }
};  