<template>
     <el-button 
        class="edit-btn"
        v-if="editing && showAddButton" 
        type="primary" 
        size="small" 
        circle 
        @click="$emit('addClick')"><el-icon><Plus /></el-icon></el-button>
     <el-button type="danger"
        class="edit-btn"
        v-if="editing && showRemoveButton" 
        size="small" 
        circle 
        :icon="Minus"
         @click="$emit('removeClick')"><el-icon><Minus /></el-icon></el-button>
</template>

<script>
import { Plus, Minus } from '@element-plus/icons-vue'
export default {
    props: {
        'showRemoveButton': {
            default: true,
            type: Boolean
        },
         'showAddButton': {
            default: true,
            type: Boolean
         },
         'editing': {
            default: false,
            type: Boolean
         }
    }
}
</script>

<style>

</style>