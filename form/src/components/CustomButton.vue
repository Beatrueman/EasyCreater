<template>
    <el-button class="btn"
                type="success" 
                plain
                :class="customClass"
                >
        <slot></slot>
    </el-button>
</template>

<script>
export default {
    props: {
        btnType: {
            type: String,
            default: ""
        }
    },
    computed: {
        customClass() {
            return 'btn-' + this.btnType;
        }

    }
}
</script>

<style>

</style>