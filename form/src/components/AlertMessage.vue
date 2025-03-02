<template>
    <el-alert 
      v-if="visible"
      :title="message"
      :type="type"
      show-icon 
      closable
      @close="closeAlert"
      class="alert-position"
    />
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps, defineEmits } from 'vue';
  
  const props = defineProps({
    message: {
      type: String,
      required: true
    },
    type: {
      type: String,
      default: 'error' // 默认为 error 类型
    },
    visible: {
      type: Boolean,
      default: false
    }
  });
  
  const emit = defineEmits(['update:visible']);
  
  const closeAlert = () => {
    emit('update:visible', false);  // 当关闭时更新父组件的 visible 状态
  };
  </script>
  
  <style scoped>
  .alert-position {
    position: fixed;
    top: 20px;
    right: 20px;
    z-index: 1000;
  }
  </style>
  