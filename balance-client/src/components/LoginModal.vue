<template>
    <div class="input-wrapper">
        <div class="input-container">
            <input type="text" v-model="username">
            <input type="text" v-model="password">
            <button @click="loginHandle">Login</button>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';
import api from '../api';
import { useAuthStore } from '../store/authStore';
import { useRouter } from 'vue-router';
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
    
}
</style>