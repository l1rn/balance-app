<template>
    <div class="table-wrapper">
        <div class="table-content">
            <table>
                <tr class="tr-header">
                    <td>ID</td>
                    <td>Username</td>
                    <td>Role</td>
                    <td>Balance</td>
                    <td>Actions</td>
                </tr>
                <tr 
                v-for="u in users"
                :key="u.id">
                    <td>{{ u.id }}</td>
                    <td>{{ u.username }}</td>
                    <td>{{ u.role }}</td>
                    <td>{{ u.balance }}</td>
                    <td class="action-cell">
                        <ButtonBase :icon="addIcon"/>
                        <ButtonBase :icon="removeIcon" />
                        <ButtonBase 
                        @click="$router.push(`/users/${u.id}`)" 
                        :icon="userIcon"/>
                    </td>
                </tr>
            </table>
        </div>
        <div class="button-content">
            <ButtonBase :icon="addIcon"/>
        </div>
    </div>
</template>

<script setup>
import ButtonBase from './ButtonBase.vue';
import { onMounted, ref } from 'vue';
import api from '../common/api';
import { addIcon, removeIcon, userIcon } from '../main';

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

<style scoped>
.table-wrapper {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    gap:0.5rem;
}

.table-content {
    display: flex;
    justify-content: center;
    width: 100%;
}

.button-content {
    width: 80%;
}

table {
    border-spacing: 0;
    border-collapse: separate;
    width: 80%;
}

.tr-header {
    font-weight: 600;
    border: 0;
    outline: 0;
}

tr td {
    padding: 0.5rem 1rem;
    border: 0;
    background: #2a2a2a;
    text-align: center;
}

.action-cell {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.25rem;
}
tr:first-child td:first-child {
    border-radius: 8px 0 0 0;
}
tr:first-child td:last-child {
    border-radius: 0 8px 0 0;
}

tr:last-child td:first-child {
    border-radius: 0 0 0 8px;
}
tr:last-child td:last-child {
    border-radius: 0 0 8px 0;
}
</style>