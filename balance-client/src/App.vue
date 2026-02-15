<template>
  <div id="app">
    <MainHeader />
    <RouterView />
  </div>
</template>
<script setup>
import { onMounted } from 'vue';
import { useAuthStore } from './store/authStore';
import MainHeader from './components/MainHeader.vue';

const authStore = useAuthStore();
onMounted(async() => {
  
})
</script>

<style scoped>

</style>
