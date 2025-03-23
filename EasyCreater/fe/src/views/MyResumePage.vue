<template>
    <main class="container">
        <div class="left-col">
          <el-tabs 
            style="margin-left: 10px;" 
            v-model="activeNameLeft" 
            class="demo-tabs"
            >
            <el-tab-pane label="我创建的" name="my-resume">
              <div class="resume-list">
                <template v-if="myResumes.length > 0">
                  <el-descriptions
                    v-for="resume in paginatedResumes"
                    :key="resume.resume_id"
                    class="resume-item"
                    :column="3"
                  >
                      <el-descriptions-item >
                        <img 
                        :src="resume.thumbnailUrl" 
                        class="template-preview-img" 
                        alt="Resume Thumbnail"
                        @click="handleClick(resume)"
                          />
                      </el-descriptions-item>

                      <el-descriptions-item label="简历名称：">
                        {{ resume.resume_name }}
                      </el-descriptions-item>
                      <el-descriptions-item label="来自模板：">
                        {{ resume.template_name }}
                      </el-descriptions-item>
                      <el-descriptions-item label="更新时间：">
                        {{ formatDate(resume.Timestamp) }}
                      </el-descriptions-item>
                      <el-descriptions-item>
                        <el-button class="btn"
                          type="danger" 
                          primary
                          size="small" 
                          :icon="Delete"
                          @click="handleDelete(resume.resume_id)"
                          >删除</el-button>
                          <el-button class="btn"
                          type="primary"
                          primary
                          size="small" 
                          :icon="Share"
                          @click="handleShare(resume.resume_id, 'share')"
                          >分享</el-button>
                      </el-descriptions-item>
                  </el-descriptions>
                </template>

                <template v-else>
                  <div class="empty-state">
                    <h4 style="margin-left: 10px;">还没有简历哦~</h4>
                    <el-button
                    :icon="Promotion"
                    style="width: 150px;"
                    @click="goToTemplate" 
                    type="success" 
                    round 
                    size="large">去制作</el-button>
                  </div>
                </template>

                <div class="pagination-container">
                  <ChangePage
                  :total="myResumes.length"
                  :pageSize="pageSize"
                  :currentPage="currentPage"
                  @pageChange="handlePageChange"
                  />
                </div>

              </div>
            </el-tab-pane>

            <el-tab-pane label="我发布的" name="my-share">
              <div class="resume-list">
                <template v-if="mySharedResumes.length > 0">
                  <el-descriptions
                    v-for="resume in paginatedSharedResumes"
                    :key="resume.resume_id"
                    class="resume-item"
                    :column="3"
                  >
                      <el-descriptions-item >
                          <img
                          :src="resume.thumbnailUrl" 
                          class="template-preview-img" 
                          alt="Resume Thumbnail"
                            @click="handleClick(resume)"
                          />
                      </el-descriptions-item>

                      <el-descriptions-item label="简历名称：">
                        {{ resume.resume_name }}
                      </el-descriptions-item>
                      <el-descriptions-item label="来自模板：">
                        {{ resume.template_name }}
                      </el-descriptions-item>
                      <el-descriptions-item label="更新时间：">
                        {{ formatDate(resume.Timestamp) }}
                      </el-descriptions-item>
                      <el-descriptions-item>
                        <el-button class="btn"
                          type="danger" 
                          primary
                          size="small" 
                          @click="handleShare(resume.resume_id, 'unshare')"
                          >取消分享</el-button>

                      </el-descriptions-item>
                  </el-descriptions>
                </template>

                <template v-else>
                  <div class="empty-state">
                    <h4 style="margin-left: 10px;">还没有分享简历哦~</h4>
                  </div>
                </template>

                <div class="pagination-container">
                  <ChangePage
                  :total="mySharedResumes.length"
                  :pageSize="pageSize"
                  :currentPage="currentPage"
                  @pageChange="handlePageChange"
                  />
                </div>

              </div>
            </el-tab-pane>

            <el-tab-pane label="我导入的" name="my-load">
              <div class="resume-list">
                <template v-if="myLoadedResumes.length > 0">
                  <el-descriptions
                    v-for="resume in paginatedLoadedResumes"
                    :key="resume.resume_id"
                    class="resume-item"
                    :column="1"
                  >
                    
                      <el-descriptions-item label="已上传的简历名称：">
                        <el-tooltip content="点击下载" placement="top">
                          <a :href="resume.url" target="_blank" style="color: #409EFF; text-decoration: none;">
                            {{ resume.resume_name }}
                          </a>
                        </el-tooltip>
                      </el-descriptions-item>
                      <el-descriptions-item label="更新时间：">
                        {{ formatDate(resume.Timestamp) }}
                      </el-descriptions-item>
                      <el-descriptions-item>
                        <el-button class="btn"
                          type="danger" 
                          primary
                          size="small" 
                          @click="handleLoadedDelete(resume.resume_id)"
                          >删除</el-button>

                      </el-descriptions-item>
                  </el-descriptions>
                </template>

                <template v-else>
                  <div class="empty-state">
                    <h4 style="margin-left: 10px;">还没有导入的简历哦~</h4>
                  </div>
                </template>

                <div class="pagination-container">
                  <ChangePage
                  :total="mySharedResumes.length"
                  :pageSize="pageSize"
                  :currentPage="currentPage"
                  @pageChange="handlePageChange"
                  />
                </div>

              </div>
            </el-tab-pane>
          </el-tabs>
          </div>

        <div class="right-col">
        <el-tabs style="margin-left: 10px;" v-model="activeNameRight" class="demo-tabs">
        <el-tab-pane label="从模板新建" name="from-template">
          <!-- 模板列表容器 -->
          <div class="template-list">
            <div class="template-item-container">
              <div class="template-preview" @click="goToMakePage(0)">
                <img src="/public/template1.jpg" class="template-preview-img">
              </div>
              <div class="template-title">
                <span>简洁专业简历模板（双栏）</span>
              </div>
            </div>

            <div class="template-item-container">
              <div class="template-preview" @click="goToMakePageSecond(0)">
                <img src="/public/template2.jpg" class="template-preview-img">
              </div>
              <div class="template-title">
                <span>美观整齐简历模板（双栏）</span>
              </div>
            </div>
          </div>
        </el-tab-pane>

        <div>
                <AlertMessage 
                v-model:visible="alertVisible"
                message="上传成功"
                type="success"
                style="width: 50%;"
                />
                
                <AlertMessage 
                v-model:visible="alertErrorVisible"
                :message="alertErrorMessage"
                type="error"
                style="width: 50%;"
                />
            </div>

        <el-tab-pane label="导入简历" name="load">
          <div class="template-item-container">
            <el-upload
                class="upload"
                drag
                :http-request="handleUpload"
                multiple
                :before-upload="beforeUpload"
            >
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                拖拽简历文件或<em>点击此处上传</em>
                </div>
                <template #tip>
                <div class="el-upload__tip">
                    仅支持 word/pdf/jpg/png, 大小需要小于500Kb
                </div>
                </template>
            </el-upload>
          </div>
        </el-tab-pane>
            <el-tab-pane label="AI 智能生成" name="ai">
                <template #label>
                    AI 智能生成
                    <el-badge value="推荐" class="item-badge" />
                </template>
                <div class="template-item-container">
                    <el-form :model="form" :rules="rules" ref="formRef" lable-length="80">
                    <el-form-item prop="position" label="目标职位">                     
                    <el-input 
                        :prefix-icon="User"
                        v-model="form.position" 
                        style="width: 480px;"
                        placeholder="请输入目标职位" />
                    </el-form-item>

                    <el-form-item prop="skills" label="技术特长">                     
                    <el-input 
                        :prefix-icon="Medal"
                        v-model="form.skills" 
                        style="width: 480px;"
                        placeholder="请输入您的特长" />
                    </el-form-item>
                    
                    <el-form-item prop="describtion" label="个人描述">                     
                        <el-input
                            v-model="form.describtion"
                            maxlength="512"
                            style="width: 480px;"
                            placeholder="请输入您的个人描述，越详细越好哦~"
                            show-word-limit
                            type="textarea"
                        />
                    </el-form-item>
                    
                    <el-form-item>
                        <AiResume 
                        :resume_data="`${form.position} ${form.skills} ${form.describtion}`"
                        :btn_text="'AI生成'"
                        />
                    </el-form-item>
                </el-form>
                </div>
            </el-tab-pane>
            </el-tabs>
        </div>
    </main>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { UploadFilled } from '@element-plus/icons-vue'
import { User, Medal, Delete, Promotion, Share } from '@element-plus/icons-vue'
import AiResume from '../components/AiResume.vue'
import { getLoadedResumeURL, deleteLoadedResume, getLoadedResumes, uploadResumeFile, fetchResumeList, deleteResume, shareResume, getSharedResume, getThumbnail } from '../apis/api'
import ChangePage from '../components/ChangePage.vue'
import AlertMessage from '../components/AlertMessage.vue'

interface ResumeData {
  resume_id: number;
  username: string;
  user_id: number;
  template_name: string;
  resume_data: string;
  Timestamp: string;
  IsShared: boolean;
  thumbnailUrl?: string;
  resume_name: string;
}

interface UploadParams {
  file: File;
}

interface LoadedResumeData {
  resume_name: string;
  resume_id: number;
  Timestamp: string;
  url: string;
}


const activeNameLeft = ref('my-resume')
const activeNameRight = ref('from-template')
const router = useRouter()
const myResumes = ref<ResumeData[]>([]);
const mySharedResumes = ref<ResumeData[]>([])
const myLoadedResumes = ref<LoadedResumeData[]>([])
const currentPage = ref(1);
const pageSize = ref(2);
const alertVisible = ref(false);
const alertErrorVisible = ref(false);
const alertErrorMessage = ref('');

const goToTemplate = () => {
    router.push('/home/template')
  };

const paginatedResumes = computed(() => {
  if (!myResumes.value) return []; 
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return myResumes.value.slice(startIndex, endIndex);
});

const paginatedSharedResumes = computed(() => {
  if (!mySharedResumes.value) return []; 
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return mySharedResumes.value.slice(startIndex, endIndex);
});

const paginatedLoadedResumes = computed(() => {
  if (!myLoadedResumes.value) return []; 
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return myLoadedResumes.value.slice(startIndex, endIndex);
});

const handlePageChange = (page: number) => {
  currentPage.value = page;
};

const fetchResumes = async () => {
  try {
    const res = await fetchResumeList()
    if (res && Array.isArray(res)) {
      // 获取所有简历的缩略图
      for (const resume of res) {
        const thumbnailUrl = await getThumbnail(resume.resume_id);
        resume.thumbnailUrl = thumbnailUrl;  // 将缩略图URL赋值到简历数据
      }
      myResumes.value = res || [];
    } 
  } catch(error) {
    console.error('Error fetching resumes:', error)
  }
  
}
// 格式化时间
const formatDate = (timestamp: string | Date): string => {
  return new Date(timestamp).toLocaleString();
};

const form = ref({
    position: "",
    skills: "",
    describtion: "",
  })

  const rules = {
    position: [{ required: true, message: '请输入您的目标职位', trigger: 'blur' }],
    skills: [{ required: true, message: '请输入您的特长', trigger: 'blur' }],
    describtion: [{ required: true, message: '请输入您的个人描述', trigger: 'blur' }],
  }

  const handleClick = (resume: { resume_id: number; template_name: string }) => {

    const resumeId = resume.resume_id;
    const templateName = resume.template_name;
    if (templateName === 'template1') {
      goToMakePage(resumeId);
    } else if (templateName === 'template2') {
      goToMakePageSecond(resumeId);
    } else {
      console.error('Invalid template name:', templateName);
    }
  }

 const goToMakePage = (resumeId: number) => {
  router.push({name: 'MakePage', query: {resume_id: resumeId}})
}

const goToMakePageSecond = (resumeId: number) => {
  router.push({name: 'MakePageSecond', query: {resume_id: resumeId}})
}

const handleDelete = (resumeId: number) => {
  try {
    const result = deleteResume(resumeId);
    if (!result) {
      console.error('Failed to delete resume.');
      return;
    }
    // 刷新页面
    myResumes.value = myResumes.value.filter(resume => resume.resume_id !== resumeId);
    mySharedResumes.value = mySharedResumes.value.filter(resume => resume.resume_id !== resumeId);
  } catch(error) {
    console.error('Error deleting resume:', error);
  }
}

const handleShare =  async (resumeId: number, isShare: string) => {
  try {
    const result = await shareResume(resumeId, isShare);
    if (!result) {
      console.error('Failed to share resume.');
      return;
    }
    mySharedResumes.value = mySharedResumes.value.filter(resume => resume.resume_id !== resumeId);
    await fetchSharedResumes();
  } catch(error) {
      console.error('Error sharing resume:', error);
  }
}

const handleLoadedDelete = async (resumeId: number) => {
  try {
    const result = await deleteLoadedResume(resumeId);
    if (!result) {
      console.error('Failed to delete resume.');
      return;
    }
    // 刷新页面
    myLoadedResumes.value = myLoadedResumes.value.filter(resume => resume.resume_id !== resumeId);
  } catch(error) {
    console.error('Error deleting resume:',error);
  }
}

const fetchSharedResumes = async () => {
  try {
    const res = await getSharedResume()
    if (res && Array.isArray(res)) {
      // 获取所有简历的缩略图
      for (const resume of res) {
        const thumbnailUrl = await getThumbnail(resume.resume_id);
        resume.thumbnailUrl = thumbnailUrl;  // 将缩略图URL赋值到简历数据
      }
      mySharedResumes.value = res;
    } 
  } catch(error) {
    console.error('Error fetching resumes:', error)
  }
}

const handleUpload = async (params: UploadParams) => {
  const file = params.file;
  const formData = new FormData();
  formData.append('file', file);

  try {
    const res = await uploadResumeFile(file);
    if (res.url) {
      alertVisible.value = true;
      setTimeout(() => { alertVisible.value = false; fetchLoadedResumes(); }, 3000);
    } else {
      throw new Error('Failed to upload resume file.');
    }
  } catch (error) {
    console.error('Error uploading resume file:', error);
    alertVisible.value = true;
    alertErrorMessage.value = '上传文件失败！请重试'
    setTimeout(() => { alertErrorVisible.value = false; }, 3000);
  }
};

const beforeUpload = (file: File) => {
  const isType = ["application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "image/jpeg", "image/png"].includes(file.type);
  const is500KB = file.size / 1024 <= 500;

  if (!isType) {
    alertErrorVisible.value = true;
    alertErrorMessage.value = '请上传PDF、Word、JPEG或PNG格式的文件，且文件大小不超过500KB'
    setTimeout(() => { alertErrorVisible.value = false; }, 3000);
    return false;
  } else if (!is500KB) {
    alertErrorVisible.value = true;
    alertErrorMessage.value = '文件大小不能超过500KB'
    setTimeout(() => { alertErrorVisible.value = false; }, 3000);
    return false;
  }
  return true
};

const fetchLoadedResumes = async () => {
  try {
    const res = await getLoadedResumes()
    if (res && Array.isArray(res)) {
      for (const resume of res) {
        const Url = await getLoadedResumeURL(resume.resume_id);
        resume.url = Url;  
      }
      myLoadedResumes.value = res || [];
    } 
  } catch(error) {
    console.error('Error fetching LoadedResumes:', error)
  } 
}

onMounted(() => {
  fetchResumes();
  fetchSharedResumes();
  fetchLoadedResumes();
});
</script>

<style scoped>
.pagination-container {
  position: absolute;
  bottom: 0;
  margin-bottom: 40px;
}

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


.el-tabs {
  width: 100%;
}
.el-tabs__item {
  font-size: 18px;
  font-weight: bold;
  color: #333;
}

.left-col {
  position: relative;
  display: flex;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
  width: 40%;
  float: left;
  height: 65vh;
  margin-left: 25px;

}

.right-col {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    width: 55%;
    float: right;
    height: 100vh;
    margin-right: 10px;
}

@media (min-width: 600px) {
    .left-col, .right-col {
        height: 85vh;
    }
}

.resume-item {
  margin-bottom: 20px;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
  background-color: white;
}

.resume-preview-img {
  margin: 5px;
  width: 60px;
  height: 60px;
  cursor: pointer;
  border-radius: 8px;
  transition: transform 0.2s ease-in-out;
}

.resume-preview-img:hover {
  transform: scale(1.1);
}

.resume-list {
  background-color: white;
}
</style>