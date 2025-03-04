<template>
    <el-pagination
                    background
                    layout="prev, pager, next"
                    :total="total"
                    :page-size="pageSize"
                    :current-page="currentPage"
                    @current-change="handlePageChange"
                    style="margin-top: 20px;"
                    >
                    </el-pagination>
</template>

<script setup>
import {defineProps, defineEmits} from 'vue'

const props = defineProps({
    total: {
        type: Number,
        required: true,
    },
    pageSize: {
        type: Number,
        required: true,
    }, 
    currentPage: {
        type: Number,
        required: true,
    }
});
const emit = defineEmits(['pageChange'])

const handlePageChange = (page) => {
    emit('pageChange', Number(page))
}

</script>