<template>
    <div>
        <ul>
            <li v-for="u in users">
                {{ u.username }} {{ u.password }} {{ u.role }} {{ u.balance }}
            </li>
        </ul>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import api from '../api';

const users = ref({})
const handleAllUsers = async() => {
    try {
        const response = await api.get("/admin/all-users")
        users.value = response.data
        console.log(response.data)
    } catch(e) {
        console.error(e)
    }
}
onMounted(() => handleAllUsers())
</script>

<style>

</style>