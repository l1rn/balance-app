<template>
  <MainHeader />
  <RouterView />
</template>
<script setup>
import { onMounted } from 'vue';
import { useAuthStore } from './store/authStore';
import api from './api';
import MainHeader from './components/MainHeader.vue';


const authStore = useAuthStore();
onMounted(async() => {
  
})
</script>

<style scoped>

</style>
