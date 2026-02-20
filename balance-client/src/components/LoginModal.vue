<template>
    <div class="input-wrapper">
        <h1>Authentication</h1>
        <div class="input-container">
            <div class="input-group">
                <div>
                    Username
                </div>
                <InputBase 
                type="text" 
                placeholder="username" 
                v-model="username" />
            </div>
            <div class="input-group">
                <div>
                    Password
                </div>
                <InputBase 
                type="password" 
                placeholder="password" 
                v-model="password" />
            </div>
            <ButtonBase @click="loginHandle" title="Login"/>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';
import api from '../common/api';
import { useAuthStore } from '../store/authStore';
import { router } from '../common/router';
import InputBase from './InputBase.vue';
import ButtonBase from './ButtonBase.vue';

const username = ref("")
const password = ref("")

const authStore = useAuthStore();

const loginHandle = async() => {
    try {
        await api.post("/auth/login", {
            "username": username.value,
            "password": password.value,
        });
        authStore.setAuthorization(true)
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
    width: 100%;
}
.input-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 80%;
    background: #2c2c2c;
    border-radius: 8px;
    padding: 1.5rem;
}

.input-group {
    display: flex;
    flex-direction: column;
    gap: 0.4rem;
}

.input-group div {
    font-size: 18px;
    text-indent: 0.1rem;
    font-weight: bold;
}
</style>