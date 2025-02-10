<template>
  <div class="flex gap-2">
    <el-tag
      v-for="(word, index) in words"
      :key="index"
      type="warning"
      effect="dark"
      size="large"
    >
      {{ word }}
    </el-tag>
  </div>
</template>

<script setup>
const props = defineProps({
  words: {
    type: Array,
    required: true,  
  }
})

</script>

<style scoped>
.el-tag {
  margin: 10px;
  font-weight: bold;
}
</style>