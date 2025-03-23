<template>
    <label>
        <input type="color" 
                class="color-input" 
                @input="$emit('color-changed', $event.target.value)"
                :value="defaultColor"
                >
        <span style="color: black;">{{ label }}</span>
    </label>
</template>

<script>
export default {
    props: {
        label: {
            type: String,
            default: '颜色选择'
        },
        defaultColor: {
            type: String,
            default: '#ffffff'
        },
        
    }
}
</script>

<style>
.label {
    display: flex;
    align-items: center;
} 

.color-input {
    margin-right: 10px;
    width: 26px;
    height: 30px;
    padding: 0;
    border: none;
    background: transparent;
    border-radius: 30px;
}
</style>