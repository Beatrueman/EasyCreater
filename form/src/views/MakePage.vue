<template>
  <el-page-header icon="">
    <template #content>
      <span class="text-large font-600 mr-3">简历生成器</span>
    </template>
  </el-page-header>

  <el-row :gutter="20" style="display: flex; align-items: start;">
    <!-- 左侧：输入框始终显示 -->
    <el-col :span="6">
      <el-collapse v-model="activePanels">
        <el-collapse-item title="基本信息" name="personalInfo">
          <el-form label-width="100px">
            <el-form-item label="姓名">
              <el-input v-model="personalInfo.name" placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="年龄">
              <el-input v-model="personalInfo.age" placeholder="请输入年龄" />
            </el-form-item>
          </el-form>
          <div class="switch-item">
            <span>显示基本信息</span>
            <el-switch v-model="showPersonalInfo" />
          </div>
        </el-collapse-item>
        <el-collapse-item title="教育背景" name="educationBackground">
          <el-form label-width="100px">
            <el-form-item label="学校">
              <el-input v-model="educationBackground.school" placeholder="请输入学校名称" />
            </el-form-item>
            <el-form-item label="学历">
              <el-input v-model="educationBackground.degree" placeholder="请输入学历" />
            </el-form-item>
          </el-form>
          <div class="switch-item">
            <span>显示教育背景</span>
            <el-switch v-model="showEducationBackground" />
          </div>
        </el-collapse-item>
        <el-collapse-item title="工作经历" name="workExperience">
          <el-form label-width="100px">
            <el-form-item label="公司">
              <el-input v-model="workExperience.company" placeholder="请输入公司名称" />
            </el-form-item>
            <el-form-item label="职位">
              <el-input v-model="workExperience.position" placeholder="请输入职位" />
            </el-form-item>
          </el-form>
          <div class="switch-item">
            <span>显示工作经历</span>
            <el-switch v-model="showWorkExperience" />
          </div>
        </el-collapse-item>
      </el-collapse>
    </el-col>

    <!-- 右侧：简历预览 -->
    <el-col :span="18">
      <div class="resume-preview">
        <h2>简历预览</h2>
        <!-- 基本信息 -->
        <div v-if="showPersonalInfo">
          <h3>基本信息</h3>
          <p>姓名：{{ personalInfo.name || "未填写" }}</p>
          <p>年龄：{{ personalInfo.age || "未填写" }}</p>
        </div>
        <!-- 教育背景 -->
        <div v-if="showEducationBackground">
          <h3>教育背景</h3>
          <p>学校：{{ educationBackground.school || "未填写" }}</p>
          <p>学历：{{ educationBackground.degree || "未填写" }}</p>
        </div>
        <!-- 工作经历 -->
        <div v-if="showWorkExperience">
          <h3>工作经历</h3>
          <p>公司：{{ workExperience.company || "未填写" }}</p>
          <p>职位：{{ workExperience.position || "未填写" }}</p>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { ref } from "vue";

// 折叠面板激活状态
const activePanels = ref(["personalInfo"]);

// 模块显示开关
const showPersonalInfo = ref(true);
const showEducationBackground = ref(true);
const showWorkExperience = ref(true);

// 编辑内容
const personalInfo = ref({
  name: "",
  age: "",
});
const educationBackground = ref({
  school: "",
  degree: "",
});
const workExperience = ref({
  company: "",
  position: "",
});
</script>

<style scoped>
.resume-preview {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}
.resume-preview h2 {
  margin-bottom: 20px;
  font-size: 20px;
}
.resume-preview h3 {
  margin-top: 20px;
  font-size: 18px;
  color: #333;
}
.resume-preview p {
  margin: 5px 0;
  color: #555;
}
.switch-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;
}
</style>
