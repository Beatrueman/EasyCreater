<template>
    <el-button
    type="primary"
    
    round
    @click="openDrawer"
    :loading="loading"
    >
    <el-icon><MagicStick /></el-icon> {{ btn_text}}
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
  import { getAIStreamBase, submitResumeData } from '../apis/api';
  import { useRouter } from "vue-router";
  import MarkdownIt from "vue3-markdown-it";
  
  export default defineComponent({
    components: { MagicStick, MarkdownIt },
    props: {
      resume_data: {
        type: String,
        required: true,
      },
      btn_text: {
        type: String,
        required: true,
      },
    },

    setup(props) {
      const aiResponse = ref("");
      const loading = ref(false);
      const drawer = ref(false);
      const router = useRouter();
  
      const openDrawer = async () => {
        loading.value = true;
        drawer.value = true;
        aiResponse.value = ""; 
  
        try {
          const taskId = await submitResumeData(props.resume_data);
          if (!taskId) {
            throw new Error("未能获取有效 task_id");
          }
          

          // 第二步：通过 SSE 接收流式响应
          const eventSource = getAIStreamBase(taskId);
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
  
      return {
        aiResponse,
        loading,
        drawer,
        openDrawer,
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
  