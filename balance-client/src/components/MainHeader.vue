<template>
    <div class="header-wrapper">
        <div class="header-container">
            <h1 @click="$router.push('/main')">Balance Handler 💸</h1>
        </div>
        <div class="button-container">
            <ButtonBase 
            title="Quit"
            @click="logoutHandle"/>
        </div>
    </div>
</template>

<script setup>
import api from '../common/api';
import { useAuthStore } from '../store/authStore';
import ButtonBase from './ButtonBase.vue';

const authStore = useAuthStore()
const logoutHandle = () => {
    document.cookie = "token=null; max-age=0; path=/;"
}
</script>

<style scoped>
.header-wrapper {
    height: 80px;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
}

.button-container {
    position: absolute;
    right: 10px;
}

h1 {
    cursor: pointer;
}
</style>