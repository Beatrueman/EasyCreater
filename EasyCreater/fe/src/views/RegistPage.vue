<template>
  <div class="background-container">
    <img src="../assets/bg.png" class="background-image">  
    <div class="common-layout">
      <el-container>
        <el-header class="header">
          <img src="../assets/logo.svg" class="logo">  
          <el-text class="title" size="large">EasyCreater</el-text>
          <el-divider class="header-divider" content-position="left"><el-text class="subtitle">一站式简历生成器</el-text></el-divider>
        </el-header>
        <el-main>
          <div class="login-container">
            <div class="login-box">
              <el-alert 
                v-if="alertVisible" 
                title="注册成功！" 
                type="success" 
                show-icon 
                closable
                @close="alertVisible = false" 
                class="alert-position" />
              <el-alert 
                v-if="alertErrorVisible" 
                :title="alertErrorMessage"
                type="error" 
                show-icon 
                closable
                @close="alertErrorVisible = false" 
                class="alert-position" />
              <el-button round class="back-button" @click="goBack"><el-icon><Back /></el-icon></el-button>
              <h2>欢迎注册</h2>
              <el-divider />
              <el-form :model="form" :rules="rules" ref="formRef" @submit.prevent="submitForm">
                <el-form-item prop="username">
                  <el-input 
                    :prefix-icon="User"
                    v-model="form.username" 
                    placeholder="请输入用户名" />
                </el-form-item>
                <el-form-item prop="password">
                  <el-input
                    :prefix-icon="Lock"
                    show-password 
                    v-model="form.password" 
                    placeholder="请输入密码" />
                </el-form-item>
                <el-form-item prop="email">
                  <el-input
                    :prefix-icon="Message"
                    v-model="form.email" 
                    placeholder="请输入邮箱" />
                </el-form-item>
                <el-form-item prop="phone">
                  <el-input
                    :prefix-icon="Cellphone"
                    v-model="form.phone" 
                    placeholder="请输入电话" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" round native-type="submit">立即注册</el-button>
                </el-form-item>
              </el-form>
            </div>
          </div>
        </el-main>
        <el-footer div class="footer">
          <el-divider />
          <el-text>Copyright © 2025 <el-text><a href="https://blog.yiiong.top">Yiiong</a></el-text>. All rights reserved.</el-text>
        </el-footer>
      </el-container>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue'
  import { User, Lock, Message, Cellphone } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  import type { ElForm } from 'element-plus'
  import { Register } from '../apis/api'

  import { AxiosError } from 'axios'

  const router = useRouter() 
  
  const form = ref({
    username: '',
    password: '',
    email: '',
    phone: ''
  })

  const rules = {
    password: [{ required: true, message: '密码不能为空', trigger: 'blur' }],
    username: [{ required: true, message: '用户名不能为空', trigger: 'blur' }],
    email: [{ required: true, message: '邮箱不能为空', trigger: 'blur' }],
    phone: [{ required: true, message: '电话不能为空', trigger: 'blur' }],
  }

  // 声明 formRef
  const formRef = ref<InstanceType<typeof ElForm> | null>(null);


  const submitForm = async () => {
  try {
    // 执行表单验证
    await formRef.value?.validate();

    // 调用注册API
    await Register(form.value);
    alertVisible.value = true;  // 显示成功提示框

    setTimeout(() => {
    alertVisible.value = true; 
  }, 2000);

    setTimeout(() => {
      router.push('/');
    }, 2000);

  } catch (error) {
    console.error('注册失败:', error);
    if (error instanceof AxiosError && error.response && error.response.data && error.response.data.msg) {
      alertErrorMessage.value = error.response.data.msg; 
    } else {
      alertErrorMessage.value = "注册失败，请重试！";
    }
    alertErrorVisible.value = true;

    setTimeout(() => {
    alertErrorVisible.value = false; 
  }, 2000);
  }
}


const alertVisible = ref(false);
const alertErrorVisible = ref(false);
const alertErrorMessage = ref('');


  const goBack = () => {
    router.push('/')
  }
</script>
    <style scoped>
    
    .background-container {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%; 
      background-color: #ffffff; /* 设置背景颜色 */;
      justify-content: center;
      align-items: center;
      overflow: hidden; 
    }
    
    
    .background-image {
      width: 100%;
      height: 100%; 
      object-fit: cover; 
    }
    
    
    
    .login-container {
      display: flex;
      justify-content: center;
      align-items: center;
      position: absolute;
      top: 50%;
      right: 200px; 
      transform: translateY(-50%);
      background-color: #ffffff;
      height: 50vh;
      border-radius: 15px;
      box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
    }
    
    .login-box {
      padding: 30px;
      width: 350px;
    }
    
    .el-form-item {
      margin-bottom: 20px;
    }
    
    .el-button {
      margin-right: 10px;
      transition: background-color 0.3s ease;
    }
    
    .footer {
      position: absolute;
      bottom: 0; 
      left: 0;
      right: 0;
      height: 80px;
      text-align: center;
    }
    
    .header {
      position: absolute;
      top: 0; 
      left: 0;
      right: 0;
      height: 80px;
    }
    
    .logo {
      width: 80px;
      position: absolute;
      left: 40px;
      top: 15px;
    }
    
    .title {
      position: absolute;
      font-weight: 200px;
      font-size: 50px;
      left: 130px;
      top: 15px;
    }
    
    .header-divider {
      top: 80px;
    }
    
    .subtitle {
      font-weight: 150;
      font-size: 18px;
    
    }
    .login-box h2 {
      margin-top: 30px; /* 调整标题与分割线的间距 */
      text-align: center;
    }

    .back-button{
      position: absolute; /* 绝对定位 */
      top: 10px; /* 距离顶部 */
      left: 20px; /* 距离左侧 */
      padding: 10px; /* 按钮的内边距 */
    }

    .alert-position {
      margin-top: -40px; 
    }

    </style>
    