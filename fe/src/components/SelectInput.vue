<template>
    <el-select 
      @change="$emit('update-selection', $event)"
      v-model="value" 
      placeholder="Select" 
      style="width: 120px; margin-left: 20px;">
      <el-option
        v-for="item in options"
        :selected="defaultOption == item.value"
        :key="item.value"
        :label="item.name"
        :value="item.value"
      />
    </el-select>
  </template>
  

  <script>

 export default {
    props: {
      options: {
        type: Array,
      },
      defaultOption: {
        type: String,
      }
    },
    data() {
    return {
      value: this.defaultOption, // 将 defaultOption 赋值给 v-model
    };
  },
  watch: {
    // 如果 defaultOption 变化，更新 value
    defaultOption(newVal) {
      this.value = newVal;
    },
  },
 }



  </script>
  

<style>

</style>