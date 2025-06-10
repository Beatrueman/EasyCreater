<template>
    <el-button
      type="warning"
      plain
      style="margin-bottom: 20px; margin-top: 0px; margin-left: 530px;"
      @click="handleClick"
      :loading="loading"
    >
      <el-icon><MagicStick /></el-icon> AI润色
    </el-button>
  
    <el-drawer v-model="drawer" title="AI优化建议" :with-header="false" size="40%">
      <div v-loading="loading">
        <span v-if="loading">正在生成优化建议...</span>
        <markdown-it v-else class="response-text" :source="aiResponse"></markdown-it>
      </div>
    </el-drawer>
  </template>
  
  <script>
  import { ref, defineComponent } from "vue";
  import { ElMessage } from "element-plus";
  import { MagicStick } from "@element-plus/icons-vue";
  import { submitResumeData, getAIStream } from '../apis/api';
  import { useRouter } from "vue-router";
  import MarkdownIt from "vue3-markdown-it";

  
  export default defineComponent({
    components: { MagicStick, MarkdownIt },
    props: {
      fromTemplate: {
        type: String,
        required: true,
      },
      
    },
    emits: ["saveResumeLocal"],
    setup(props, { emit }) {
      const aiResponse = ref("");
      const loading = ref(false);
      const drawer = ref(false);
      const router = useRouter();
  
      const openDrawer = async () => {
        loading.value = true;
        drawer.value = true;
        aiResponse.value = ""; 

        const resumeData = localStorage.getItem(`resumeData_${props.fromTemplate}`);
        if (!resumeData) {
          aiResponse.value = "未找到简历数据，请先保存您的简历";
          loading.value = false;
          return;
        }
  
        try {
          const taskId = await submitResumeData(resumeData);
          if (!taskId) {
            throw new Error("未能获取有效 task_id");
          }
          

          // 第二步：通过 SSE 接收流式响应
          const eventSource = getAIStream(taskId);
          console.log(eventSource);
          
          eventSource.onmessage = (event) => {
            const data = JSON.parse(event.data);
            if (data.content) {
              aiResponse.value += data.content;
            }
          };

          eventSource.onerror = (err) => {
            console.error("SSE 错误", err);
            ElMessage.error("AI 流式请求失败");
            loading.value = false;
            eventSource.close();
          };

          eventSource.onopen = () => {
            console.log("SSE 连接已建立");
          };
        } catch (error) {
          aiResponse.value = "请求失败，请检查网络或后端状态";
          ElMessage.error("请求失败，请检查网络或后端状态");
        } finally {
          loading.value = false;
        }
      };

      const handleClick = async() => {
        emit("saveResumeLocal");
        setTimeout(() => {
          openDrawer();
        },100);
      }
  
      return {
        aiResponse,
        loading,
        drawer,
        openDrawer,
        handleClick,
      };
    },
  });
  </script>
  
  <style scoped>
  .response-text {
    font-family: Arial, sans-serif;
    line-height: 1.6;
  }
  </style>
  
