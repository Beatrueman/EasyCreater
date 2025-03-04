<template>
    <div class="sidebar">
        <slot></slot>
    </div>
</template>

<style scoped>
.sidebar {
    width: 250px;
    height: 550px;
    position: fixed;
    top: 85px;
    right: 60px;
    padding: 20px;
    
    background-color: rgba(255, 255, 255, 0.55);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgb(255, 255, 255);
}


</style>