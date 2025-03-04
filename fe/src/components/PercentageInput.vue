<template>
    <label>
        <input 
            class="custom-input"
            type="number"
            :min="min"
            :max="max"
            :value="currentValue"
            @change="$emit('percentage-changed', parseInt($event.target.value))"
            >
        % {{ label }}
    </label>
</template>

<script>
    export default {
        props: {
            label: {
                type: String
            },
            min: {
                type: Number,
                default: 0
            },
            max: {
                type: Number,
                default: 100
            },
            currentValue: {
                type: Number,
                default: 0
            }
        }
    }
</script>

<style scoped>
label {
    display: flex;
    align-items: center;
}

input {
    width: 50px;
    margin-right: 8px;
    display: block;
}

.custom-input {
    padding: 6px 10px;
    background-color: white;
    border: 1px solid #ccc;
    color: #495057;
    border-radius: 4px;
    font-size: 16px;
    line-height: 12px;
}

input:focus {
    outline-color: #82c0cc;
    border-color: #82c0cc;
}

</style>