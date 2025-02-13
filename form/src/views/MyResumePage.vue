<template>
    <main class="container">
        <div class="left-col">
            <el-tabs style="margin-left: 10px;" v-model="activeNameLeft" class="demo-tabs">
        <el-tab-pane label="我创建的" name="my-resume">
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

        <el-tab-pane label="我收藏的" name="my-favorite">
          <div class="template-item-container">
            
          </div>
        </el-tab-pane>
        <el-tab-pane label="我发布的" name="my-share">
          <div class="template-item-container">
            
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

        <el-tab-pane label="导入简历" name="load">
          <div class="template-item-container">
            <el-upload
                class="upload-demo"
                drag
                action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                multiple
            >
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                拖拽简历文件或<em>点击此处上传</em>
                </div>
                <template #tip>
                <div class="el-upload__tip">
                    jpg/png 需要小于500Kb
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
                            placeholder="请输入您的个人描述"
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
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { UploadFilled } from '@element-plus/icons-vue'
import { User, Medal } from '@element-plus/icons-vue'
import AiResume from '../components/AiResume.vue'


const activeNameLeft = ref('my-resume')
const activeNameRight = ref('from-template')
const router = useRouter()

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

.left-col {
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

</style>