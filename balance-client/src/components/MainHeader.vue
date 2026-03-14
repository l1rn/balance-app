<template>
    <div class="header-wrapper">
        <div class="header-container">
            <h1 @click="$router.push('/main')">Balance Handler 💸</h1>
        </div>
        <div
        v-if="authStore.isAuthenticated" 
        class="button-container">
            <ButtonBase 
            :icon="logoutIcon"
            @click="logoutHandle"/>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { logoutIcon } from '../main';
import { useAuthStore } from '../store/authStore';
import ButtonBase from './ButtonBase.vue';
import api from '../common/api';

const router = useRouter()
const authStore = useAuthStore()
const logoutHandle = async() => {
    await api.post("/auth/logout", { method: "POST", credentials: "include" })
    authStore.setAuthorization(false)
    window.location.assign("/login")
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