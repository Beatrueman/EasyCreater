<template>
    <AlertMessage 
    v-model:visible="alertVisible"
    message="图片上传成功！"
    type="success"
    />
    
    <AlertMessage 
    v-model:visible="alertErrorVisible"
    :message="alertErrorMessage"
    type="error"
    />

    <el-upload
        class="avatar-uploader"
        :show-file-list="false"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload"
    >
        <img v-if="imageUrl" :src="imageUrl" class="avatar" />
        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
    </el-upload>
</template>

<style scoped>
.avatar-uploader .avatar {
  width: 168px;
  height: 168px;
  display: block;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
  background-color: rgba(255, 255, 255, 0.25);
  backdrop-filter: blur(6px);
  -webkit-backdrop-filter: blur(6px);
  border: 0.8px solid rgba(255, 255, 255, 0.18);
  box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
  -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
  border-radius: 12px;
  -webkit-border-radius: 12px;
}
</style>

<script lang="ts" setup>
import { ref, defineProps, defineEmits } from 'vue';
import { uploadAvatar } from '../apis/api';

const alertVisible = ref(false);
const alertErrorVisible = ref(false);
const alertErrorMessage = ref('');

const props = defineProps({
    sizeLimit: {
        type: Number,
        default: 2
    },
});

const emit = defineEmits(['update:imageUrl']); 

const imageUrl = ref<string | null>(null);

const beforeAvatarUpload = (file: File): boolean => {
    const isImage = file.type.startsWith('image/');  // 使用 startsWith 来检查文件类型
    const isValidSize = file.size / 1024 / 1024 < props.sizeLimit;  // 限制文件大小

    if (!isImage) {
        alertErrorMessage.value = "只能上传图片文件";
        alertErrorVisible.value = true;
        return false;  
    }

    if (!isValidSize) {
        alertErrorMessage.value = "图片大小不能超过5MB";
        alertErrorVisible.value = true;
        return false; 
    }

    // 文件合法，开始进行 base64 转换
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = async () => {
        const base64Image = reader.result as string;  // 将结果转为字符串
        try {
            await uploadAvatar(base64Image);  // 上传 Base64 图片
            imageUrl.value = base64Image;  // 更新本地预览图片
            alertVisible.value = true;
            setTimeout(() => { alertVisible.value = false; }, 2000);
            emit('update:imageUrl', base64Image);  // 向父组件传递更新的图片URL
        } catch (error) {
            alertErrorMessage.value = '上传头像失败';
            alertErrorVisible.value = true;
            console.error('上传图片失败', error);
        }
    };
    return false;  // 阻止默认上传行为
};


const handleAvatarSuccess = (response: any, file: File) => {
  console.log("上传成功:", response);
};
</script>