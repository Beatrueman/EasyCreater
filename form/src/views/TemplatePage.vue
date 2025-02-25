<template>
  <main class="container">
    <div class="tab-container">
      <el-tabs style="margin-left: 10px;" v-model="activeName" class="demo-tabs">
        <el-tab-pane label="全部" name="first">
          <!-- 模板列表容器 -->
          <div class="template-list">
            <div class="template-item-container">
              <div class="template-preview" @click="goToMakePage">
                <img src="/public/template1.jpg" class="template-preview-img">
              </div>
              <div class="template-title">
                <span>简洁专业简历模板（双栏）</span>
              </div>
            </div>

            <div class="template-item-container">
              <div class="template-preview" @click="goToMakePageSecond">
                <img src="/public/template2.jpg" class="template-preview-img">
              </div>
              <div class="template-title">
                <span>美观整齐简历模板（双栏）</span>
              </div>
            </div>
          </div>
        </el-tab-pane>

        <el-tab-pane label="通用" name="second">
          <div class="template-item-container">
            <div class="template-preview" @click="goToMakePage">
              <img src="/public/template1.jpg" class="template-preview-img">
            </div>
            <div class="template-title">
              <span>简洁专业简历模板（双栏）</span>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="实习" name="third">
          <div class="template-item-container">
            <div class="template-preview" @click="goToMakePageSecond">
              <img src="/public/template2.jpg" class="template-preview-img">
            </div>
            <div class="template-title">
              <span>美观整齐简历模板（双栏）</span>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </main>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const activeName = ref('first')
const router = useRouter()

const goToMakePage = () => {
  router.push('/home/template/first')
}

const goToMakePageSecond = () => {
  router.push('/home/template/second')
}
</script>

<style scoped>
/* 模板项容器：水平排列 */
.template-item-container {
  display: flex; 
  flex-direction: column;  
  margin: 10px; 
}

.template-list {
  display: flex; /* 使用 flexbox */
  flex-wrap: wrap; /* 允许换行 */
  justify-content: flex-start; /* 开始对齐 */
  gap: 20px; /* 每个模板之间的水平和垂直间距 */
}

.template-preview {
  border: 2px solid #ddd;
  border-radius: 10px;
  padding: 10px;
  width: 150px;
  height: 160px;
  left: 0;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.template-preview:hover {
  transform: scale(1.05);
  box-shadow: 0px 4px 12px rgba(64, 158, 255, 0.4);
}

/* 图片样式 */
.template-preview-img {
  width: 105px;
  height: 155px;
  object-fit: cover;
}

/* 模板标题样式 */
.template-title {
  margin-top: 5px;
}


/* 其他样式保持不变 */
.el-tabs__item {
  font-size: 18px;
  font-weight: bold;
  color: #333;
}

.tab-container {
  box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
  height: 570px;
  margin-top: -10px;
  margin-bottom: -50px;
}

@media (min-width: 600px) {
  .tab-container {
    height: 82vh;
  }
}
</style>
