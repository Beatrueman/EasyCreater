<template>
    <el-card style="max-height: 520px; max-width: 480px; margin: 20px auto;">
      <template #header>
        <div class="card-header">
          <span>123</span>
        </div>
      </template>
      <p v-for="o in 4" :key="o" class="text item"></p>
    </el-card>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { jwtDecode } from 'jwt-decode';
  
  const username = ref<string | null>(null);
  const email = ref<string | null>(null);
  
  onMounted(() => {
    const token = localStorage.getItem('jwt-token');
    if (token) {
      try {
        const decodedToken: any = jwtDecode(token);
        username.value = decodedToken.username;
        email.value = decodedToken.email;
      } catch (error) {
        console.error('解析token失败:', error);
      }
    }
  });
  </script>
  
  <style scoped>
  .card-header {
    font-size: 18px;
    font-weight: bold;
  }
  .text {
    margin: 10px 0;
  }
  </style>
  