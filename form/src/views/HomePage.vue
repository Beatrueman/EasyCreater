<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-row :gutter="20">
        <el-col :span="4"><div class="grid-content ep-bg-purple" />
          <img src="../assets/logo.svg" class="logo">  
          <el-text class="title" size="small">EasyCreater</el-text>
        </el-col>
        <el-col :span="16"><div class="grid-content ep-bg-purple" />part2</el-col>
        <el-col :span="4"><div class="grid-content ep-bg-purple" />
          <div class="block">
          <el-avatar shape="square" :size="50" :src="squareUrl" />

  <el-dropdown :hide-on-click="false">
    <span class="el-dropdown-link">
      {{ username || '请登录' }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item divided @click="goToAbout">个人中心</el-dropdown-item>
        <el-dropdown-item divided @click="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>

        </div>
        </el-col>
        </el-row>
      </el-header>
      <el-container>
        <el-aside>
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
      >
      <el-menu-item index="1" @click="goToHome">
          <el-icon><setting /></el-icon>
          <span>首页</span>
        </el-menu-item>
        <el-menu-item index="2" @click="goToMake">
          <el-icon><setting /></el-icon>
          <span>新建简历</span>
        </el-menu-item>
        <el-sub-menu index="3">
          <template #title>
            <el-icon><location /></el-icon>
            <span>我的简历</span>
          </template>
            <el-menu-item index="3-1">我创建的</el-menu-item>
            <el-menu-item index="3-2">我收藏的</el-menu-item>
        </el-sub-menu>
        <el-menu-item index="4">
          <el-icon><location /></el-icon>
          <span>简历模板</span>
        </el-menu-item>
        <el-menu-item index="5">
          <el-icon><location /></el-icon>
          <span>简历广场</span>
        </el-menu-item>
        <el-menu-item index="6">
          <el-icon><location /></el-icon>
          <span>经验分享</span>
        </el-menu-item>
      </el-menu>
        </el-aside>
        <el-container>
          <el-main>
            <router-view />
          </el-main>
          <el-footer>
            <el-text>Copyright © 2025 <el-text><a href="https://blog.yiiong.top">Yiiong</a></el-text>. All rights reserved.</el-text>
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.el-dropdown-link{
  margin-top: 6px;
  padding: 10px;
}

@import "../style/common.scss";
.common-layout {
  height: 100%;
  width: 100%;
}

.el-container {
  height: 100%;
  width: 100%;
}

.el-header
{
  background-color: #f4f6f7;
  color: #333;
  text-align: center;
  line-height: 60px;
  width: 100%;
  height: 70px;
}

.el-footer {
  background-color: #f4f6f7;
  color: #333;
  text-align: center;
  line-height: 60px;
  width: 100%;
  height: 64px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 160px;
}

.el-aside {
  color: #333;
  background-color: #f4f6f7;
  width: 230px;
}

.el-row {
  display: flex;
  width: 100%; /* 确保el-row宽度填充父容器 */
}

.logo {
  width: 60px;
  position: absolute;
  left: 1px;
  top: 5px;
}

.title {
  position: absolute;
  font-weight: 200px;
  font-size: 35px;
  left: 63px;
  top: 5px;
}

.block {
  margin-top: 10px;
}

.el-col,
.el-sub-menu {
  width: 230px;
}


</style>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
// 使用命名导入
import { jwtDecode } from "jwt-decode";
import { reactive, toRefs } from 'vue'
import { useRouter } from 'vue-router';

const username = ref<string | null>(null);
const router = useRouter();

onMounted(() => {
  username.value = getUsernameFromToken(); // 获取用户名
});

const getUsernameFromToken = (): string | null => {
  const token = localStorage.getItem('jwt-token');
  if (token) {
    try {
      const decodedToken: any = jwtDecode(token);  // 解析token
      return decodedToken.username; // 获取用户名
    } catch (error) {
      console.error('解析token失败:', error);
      return null;
    }
  }
  return null;
};

const state = reactive({
  circleUrl:
    'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
  squareUrl:
    'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png',
  sizeList: ['small', '', 'large'] as const,
})

const { circleUrl, squareUrl, sizeList } = toRefs(state)

const logout = () => {
  localStorage.removeItem('jwt-token');
  username.value = null;
  router.push('/');
}

const goToAbout = () => {
  router.push('/home/about');  
};

const goToHome = () => {
  router.push('/home'); 
};

const goToMake = () => {
  router.push('/home/make'); 
};
</script>