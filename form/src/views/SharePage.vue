<template>
  <main class="container">
    <templplate v-if="mySharedResumes.length > 0" >
      <div class="tab-container" >
        <div class="template-list">
          <div v-for="(resume, index) in paginatedSharedResumes" :key="index" class="template-item-container">
            <div class="template-preview" @click="handleClick(resume)">
              <img :src="resume.thumbnailUrl" class="template-preview-img" alt="Resume Thumbnail">
              <div>
                <span>{{ resume.resume_name }}</span>
              </div>
              <span>{{ resume.username }}分享的简历</span>
            </div>
          </div>
        </div>
      </div>
    </templplate>

      <templplate v-else>
        <div class="empty-state">
          <h4 style="margin-left: 10px;">还没有人分享简历哦，快去分享吧~</h4>
          <el-button
          :icon="Promotion"
          style="width: 150px;"
          @click="goToMyResume" 
          type="success" 
          round 
          size="large">去制作</el-button>
        </div>
      </templplate>
    <div class="pagination-container">
          <ChangePage
          :total="mySharedResumes.length"
          :pageSize="pageSize"
          :currentPage="currentPage"
          @pageChange="handlePageChange"
          />
        </div>
  </main>
</template>

<style scoped>
.pagination-container {
  position: absolute;
  bottom: 0;
  margin-bottom: 80px;
}

.tab-container {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    height: 80vh;
    margin-top: -10px;
    margin-bottom: -5vh;
  }

@media (min-width: 600px) {
  .tab-container {
    height: 79vh;
    }
}

.template-list {
  padding: 5%;
  display: grid; 
  grid-template-columns: repeat(5, 1fr);
  justify-content: center;
  gap: 30px; /* 每个模板之间的水平和垂直间距 */
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
  justify-content: center;
}

.template-item-container {
  margin-bottom: 50px;
  width: 100%; 
  display: flex;
  flex-direction: column; /* 垂直排列内容 */
  align-items: center; /* 内容居中 */
  gap: 10px; /* 内部元素之间的间距 */
}
</style>

<script setup lang="ts">
import { getAllSharedResume, getThumbnail } from '../apis/api'
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { Promotion } from '@element-plus/icons-vue'

const router = useRouter();

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


const mySharedResumes = ref<ResumeData[]>([]);
const currentPage = ref(1);
const pageSize = ref(10);

const handleClick = (resume: { resume_id: number; template_name: string }) => {
  const resumeId = resume.resume_id;
  const templateName = resume.template_name;
  if (templateName === 'template1') {
    goToMakePage(resumeId, false);
  } else if (templateName === 'template2') {
    goToMakePageSecond(resumeId, false);
  } else {
    console.error('Invalid template name:', templateName);
  }
};

const goToMakePage = (resumeId: number, is_display: boolean) => {
  router.push({name: 'MakePage', query: {resume_id: resumeId, is_display: is_display.toString()}})
}

const goToMakePageSecond = (resumeId: number, is_display: boolean) => {
  router.push({name: 'MakePageSecond', query: {resume_id: resumeId, is_display: is_display.toString()}})
}

const goToMyResume = () => {
  router.push({ name: 'MyResumePage' });
}

const paginatedSharedResumes = computed(() => {
  if (!mySharedResumes.value) return []; 
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return mySharedResumes.value.slice(startIndex, endIndex);
});

const handlePageChange = (page: number) => {
  currentPage.value = page;
};
const fetchSharedResumes = async () => {
  try {
    const res = await getAllSharedResume();
    if (res && Array.isArray(res)) {
      // 获取所有简历的缩略图
      for (const resume of res) {
        const thumbnailUrl = await getThumbnail(resume.resume_id);
        resume.thumbnailUrl = thumbnailUrl;  // 将缩略图URL赋值到简历数据
      }
      mySharedResumes.value = res;
    }
  } catch (error) {
    console.error('Error fetching resumes:', error);
  }
};

onMounted(async () => {
  fetchSharedResumes();
});
</script>

