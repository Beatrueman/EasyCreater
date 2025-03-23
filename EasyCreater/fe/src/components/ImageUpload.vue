<template>
    <input 
        type="file" 
        accept="image/jepg, image/png, image/jpg"
        @change="changeImage"
    >
</template>

<script>
export default {
    methods: {
        changeImage(event) {
            const selectedFile = event.target.files[0];
            if(selectedFile) {
                const reader = new FileReader();

                reader.addEventListener('load', () => {
                    this.$emit('imageChanged', reader.result)
                });

                reader.readAsDataURL(selectedFile);
            
    
            }
        }
    }
}
</script>

<style>

</style>

