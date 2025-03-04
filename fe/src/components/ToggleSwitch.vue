<template>
    <el-switch class="switch-wrapper" 
            v-model="isEditing" 
            @change="onToggle"/>
    <span style="margin-left: 10px; color: black;">{{ label }}</span>
</template>

<script>
export default {
    props: {
        label: {
            type: String
        },
        defaultValue: {
            type: Boolean,
            default: false
        },
        toggleActive: {
            type: Boolean,
            default: true
        }   
    },

    data() {
        return {
            isEditing: this.toggleActive
        }
    },

    watch: {
        // 监听父组件传入的 toggleActive 变化
        toggleActive(newValue) {
            this.isEditing = newValue;
        },
        // 同步内部 isEditing 变化到父组件
        isEditing(newValue) {
            this.$emit("switchToggled", newValue);
        },
    },

    // methods: {
    //     onToggle(event) {
    //         this.$emit("switchToggled", this.isEditing)
    //     }
    // }
}
</script>

<style>

</style>