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

export const saveResume = async (resumeData: any, thumbnailDataUrl: string) => {
  try {
    const response = await axiosInstance.post(`/user/resume/save`, {
      resume_data: JSON.stringify(resumeData), 
      thumbnail: thumbnailDataUrl 
    });
    console.log(response)
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

export const uploadAvatar = async (ImgBase64String: string) => {
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

export const shareResume = async (resumeId: number, IsShare: string) => {
  try {
    const response = await axiosInstance.put(`/user/resume/share/${resumeId}`, {
      "action": IsShare,
    });
    return response;
  } catch(error) {
    console.error("分享简历失败:", error);
    throw error;
  }
};

export const getSharedResume = async () => {
  try {
    const response = await axiosInstance.get(`/user/resume/share`);
    return response.data;
  } catch(error) {
    console.error("获取该用户已分享简历失败:", error);
    throw error;
  }
};


export const getAllSharedResume = async () => {
  try {
    const response = await axiosInstance.get(`/user/resume/share?is_all=true`);
    return response.data;
  } catch(error) {
    console.error("获取该用户已分享简历失败:", error);
    throw error;
  }
};

// 上传缩略图到后端oss
export const uploadThumbnail = async (resumeId: number, file: File) => {
  try {
    const formData = new FormData();
    formData.append("file", file);
    const response = await axiosInstance.post(`/user/resume/thumbnail_upload/${resumeId}`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    return response.data;
  } catch(error) {
    console.error("上传缩略图失败",error);
  }
};

export const getThumbnail = async (resumeId: number) => {
  try {
    const response = await axiosInstance.get(`/user/resume/thumbnail/${resumeId}`);
    return response.data;
  } catch(error) {
    console.error("获取缩略图失败",error);
  }
};

// 上传简历文件
export const uploadResumeFile = async (file: File) => {
  try {
    const formData = new FormData();
    formData.append("file", file);
    const response = await axiosInstance.post(`/user/resume/upload`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    return response.data;
  } catch(error) {
    console.error("上传简历文件失败",error);
  }
}

// 获取简历文件
export const getLoadedResumes = async () => {
  try {
    const response = await axiosInstance.get(`/user/resume/get_loaded`);
    return response.data;
  } catch(error) {
    console.error("获取简历文件失败",error);
  }
}

export const deleteLoadedResume = async (resumeId: number) => {
  try {
    const response = await axiosInstance.delete(`/user/resume/delete_loaded/${resumeId}`);
    return response;
  } catch(error) {
    console.error("删除简历文件失败",error);
  }
}

export const getLoadedResumeURL = async (resumeId: number) => {
  try {
    const response = await axiosInstance.get(`/user/resume/get_loaded_url/${resumeId}`);
    return response.data;
  } catch(error) {
    console.error("获取简历文件失败",error);
  }
}

export const getIdeas = async () => {
  try {
    const response = await axiosInstance.get(`/user/resume/get_idea`);
    return response.data;
  } catch(error) {
    console.error("获取灵感词汇失败",error);
  }
}

// 切换简历点赞状态
export const toggleResumeLike = async (resumeId: number) => {
  try {
    const response = await axiosInstance.post(`/user/resume/like/${resumeId}`);
    return response;
  } catch (error) {
    console.error("切换点赞状态失败:", error);
    throw error;
  }
};

// 获取简历点赞状态
export const getResumeLikeStatus = async (resumeId: number) => {
  try {
    const response = await axiosInstance.get(`/user/resume/like/${resumeId}`);
    return response;
  } catch (error) {
    console.error("获取点赞状态失败:", error);
    throw error;
  }
};