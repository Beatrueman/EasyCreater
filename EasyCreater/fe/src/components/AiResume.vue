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
  import { askAIBase } from '../apis/api';
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
  
        try {
          const response = await askAIBase(props.resume_data);
          
          if (!response) {
            throw new Error("API 返回数据为空");
        }

          // 检查后端状态码
          if(response.status == 2005) {
            ElMessage.error("登录已过期，请重新登录");
            alert('登录已过期，请重新登录');
            localStorage.removeItem("jwt-token"); // 清除 token
            router.replace({ name: "login" }); // 跳转到登录页
            return;
          }

          aiResponse.value = response.reply || "AI优化建议生成失败";
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
  