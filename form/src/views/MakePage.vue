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
          <el-form class="form-basic" label-width="90px">
            <el-form-item label="个人照片">
              <el-upload
                  class="upload-demo"
                  action="#"
                  list-type="picture-card"
                  :on-change="handleAvatarChange"
                  :show-file-list="false"
                  :before-upload="beforeAvatarUpload"
              >
                <i v-if="personalInfo.avatar" class="el-icon-check"></i>
                <img v-if="personalInfo.avatar" :src="personalInfo.avatar" class="avatar" />
                <div>
                  <i class="el-icon-plus"></i>
                  <div>上传头像</div>
                </div>
              </el-upload>
            </el-form-item>
            <el-form-item label="姓名">
              <el-input v-model="personalInfo.name" placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="性别">
              <el-select v-model="personalInfo.gender" placeholder="请选择性别">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="学历">
              <el-input v-model="personalInfo.education" placeholder="请输入学历" />
            </el-form-item>
            <el-form-item label="年龄">
              <el-input v-model="personalInfo.age" placeholder="请输入年龄" />
            </el-form-item>
            <el-form-item label="电话">
              <el-input v-model="personalInfo.phone" placeholder="请输入电话" />
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="personalInfo.email" placeholder="请输入邮箱" />
            </el-form-item>
          </el-form>
        </el-collapse-item>

        <el-collapse-item title="教育背景" name="educationBackground">
          <el-form class="form-education" label-width="90px">
            <el-form-item label="学校名称">
              <el-input v-model="educationBackground.school" placeholder="请输入学校名称" />
            </el-form-item>
            <el-form-item label="专业名称">
              <el-input v-model="educationBackground.major" placeholder="请输入专业名称" />
            </el-form-item>
            <el-form-item label="GPA">
              <el-input v-model="educationBackground.gpa" placeholder="例如：X.X/4.0" />
            </el-form-item>
            <el-form-item label="班级排名">
              <el-input v-model="educationBackground.rank" placeholder="例如：2/35" />
            </el-form-item>
            <el-form-item label="起止时间">
              <el-date-picker
                  v-model="educationBackground.period"
                  type="daterange"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                  style="width: 100%;"
              />
            </el-form-item>
            <el-form-item label="描述">
              <el-input type="textarea" v-model="educationBackground.description" placeholder="请输入教育背景描述" />
            </el-form-item>
          </el-form>
        </el-collapse-item>

        <el-collapse-item title="工作经历" name="workExperience">
          <el-form class="form-experience" label-width="90px">
            <el-form-item label="公司名称">
              <el-input v-model="workExperience.company" placeholder="请输入公司名称" />
            </el-form-item>
            <el-form-item label="职位名称">
              <el-input v-model="workExperience.position" placeholder="请输入职位名称" />
            </el-form-item>
            <el-form-item label="起止时间">
              <el-date-picker
                  v-model="workExperience.period"
                  type="daterange"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                  style="width: 100%;"
              />
            </el-form-item>
            <el-form-item label="描述">
              <el-input type="textarea" v-model="workExperience.description" placeholder="请输入工作经历描述" />
            </el-form-item>
          </el-form>
        </el-collapse-item>

        <el-collapse-item title="专业技能" name="skills">
          <el-form label-width="90px">
            <el-form-item label="技能">
              <el-input v-model="skills" placeholder="例如：JavaScript, Vue, React" />
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </el-col>

    <!-- 右侧：简历预览 -->
    <el-col :span="18">
      <div class="resume-preview">
        <h2>简历预览</h2>
        <!-- 基本信息 -->
        <div class="section">
          <h3>基本信息</h3>
          <img v-if="personalInfo.avatar" :src="personalInfo.avatar" class="avatar-preview" />
          <p>姓名：{{ personalInfo.name || "未填写" }}</p>
          <p>性别：{{ personalInfo.gender || "未填写" }}</p>
          <p>学历：{{ personalInfo.education || "未填写" }}</p>
          <p>年龄：{{ personalInfo.age || "未填写" }}</p>
          <p>电话：{{ personalInfo.phone || "未填写" }}</p>
          <p>邮箱：{{ personalInfo.email || "未填写" }}</p>
        </div>
        <!-- 教育背景 -->
        <div class="section">
          <h3>教育背景</h3>
          <p>学校名称：{{ educationBackground.school || "未填写" }}</p>
          <p>专业名称：{{ educationBackground.major || "未填写" }}</p>
          <p>GPA：{{ educationBackground.gpa || "未填写" }}</p>
          <p>班级排名：{{ educationBackground.rank || "未填写" }}</p>
          <p>起止时间：{{ educationBackground.period.join(' ~ ') || "未填写" }}</p>
          <p>描述：{{ educationBackground.description || "未填写" }}</p>
        </div>
        <!-- 工作经历 -->
        <div class="section">
          <h3>工作经历</h3>
          <p>公司名称：{{ workExperience.company || "未填写" }}</p>
          <p>职位名称：{{ workExperience.position || "未填写" }}</p>
          <p>起止时间：{{ workExperience.period.join(' ~ ') || "未填写" }}</p>
          <p>描述：{{ workExperience.description || "未填写" }}</p>
        </div>
        <!-- 专业技能 -->
        <div class="section">
          <h3>专业技能</h3>
          <p>{{ skills || "未填写" }}</p>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { ref } from "vue";

// 折叠面板激活状态
const activePanels = ref(["personalInfo"]);

// 编辑内容
const personalInfo = ref({
  avatar: "",
  name: "",
  gender: "",
  education: "",
  age: "",
  phone: "",
  email: "",
});
const educationBackground = ref({
  school: "",
  major: "",
  gpa: "",
  rank: "",
  period: [],
  description: "",
});
const workExperience = ref({
  company: "",
  position: "",
  period: [],
  description: "",
});
const skills = ref("");

// 处理头像上传
const handleAvatarChange = (file) => {
  const reader = new FileReader();
  reader.onload = (event) => {
    personalInfo.value.avatar = event.target.result; // 设置头像
  };
  if (file.file) {
    reader.readAsDataURL(file.file);
  }
};

// 在上传之前进行验证
const beforeAvatarUpload = (file) => {
  const isJPG = file.type === 'image/jpeg' || file.type === 'image/png';
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isJPG) {
    this.$message.error('上传头像图片只能是 JPG 或 PNG 格式!');
  }
  if (!isLt2M) {
    this.$message.error('上传头像图片大小不能超过 2MB!');
  }
  return isJPG && isLt2M;
};
</script>

<style scoped>
.upload-demo {
  display: inline-block;
  width: 100%;
  height: 100px;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  text-align: center;
  line-height: 100px;
  color: #8c8c8c;
}

.avatar {
  width: 100%;
  height: 100%;
  border-radius: 6px;
}

.avatar-preview {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  margin-top: 10px;
  border: 1px solid #e1e1e1;
}
.resume-preview {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}
.resume-preview h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}
.section {
  margin-top: 20px;
  padding: 10px;
  background-color: #fff;
  border: 1px solid #e1e1e1;
  border-radius: 4px;
}
.section h3 {
  margin-top: 0;
  font-size: 20px;
  color: #555;
}
.section p {
  margin: 5px 0;
  color: #666;
}
</style>