<template>
    <input type="text" v-model="username">
    <input type="text" v-model="password">
    <button @click="loginHandle">Login</button>
</template>
<script setup>
import { ref } from 'vue';
import api from '../api';

const username = ref("")
const password = ref("")

const loginHandle = async() => {
    try {
        const response = await api.post("/auth/login", {
            "username": username.value,
            "password": password.value,
        });
        console.log("success:", (response.data))
    } catch(e){
        console.error(e)
    }
}
</script>
<style>
</style>