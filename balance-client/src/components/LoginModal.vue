<template>
    <div class="input-wrapper">
        <h1>Authentication</h1>
        <div class="input-container">
            <input type="text" v-model="username">
            <input type="text" v-model="password">
            <button @click="loginHandle">Login</button>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';
import api from '../common/api';
import { useAuthStore } from '../store/authStore';
import { router } from '../common/router';

const username = ref("")
const password = ref("")

const authStore = useAuthStore();

const loginHandle = async() => {
    try {
        const response = await api.post("/auth/login", {
            "username": username.value,
            "password": password.value,
        });
        console.log("success:", (response.data))
        authStore.confirmAuth()
        router.push("/main")
    } catch(e){
        console.error(e)
    }
}
</script>
<style scoped>
.input-wrapper {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    color: white;
}
.input-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 2rem;
    background: #2c2c2c;
    border-radius: 8px;
}
</style>