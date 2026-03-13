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
                    <template
                    v-if="u.id !== 0">
                        <td>{{ u.id }}</td>
                        <td>{{ u.username }}</td>
                        <td>{{ u.role }}</td>
                        <td>{{ u.balance }}</td>
                        <td
                        v-if="u.username.length > 0" 
                        class="action-cell">
                            <div class="action-wrapper">
                                <ButtonBase :icon="addIcon"/>
                                <ButtonBase :icon="removeIcon" />
                                <ButtonBase 
                                @click="$router.push(`/users/${u.id}`)" 
                                :icon="userIcon"/>
                            </div>
                        </td>
                    </template>
                    <template
                    v-else>
                        <td colspan="2">
                            <div class="input-container">
                                <span>
                                    <b>👤 user:</b>
                                </span>
                                <InputBase 
                                class="table-input"
                                placeholder="username" 
                                v-model="userRequest.username"/>
                                <span>
                                    <b>🔒 password:</b>
                                </span>
                                <InputBase
                                class="table-input"
                                placeholder="password" 
                                v-model="userRequest.password"/>
                            </div>
                        </td>
                        <td colspan="2">
                            <div class="input-container">
                                <span>
                                    <b>🧩 role:</b>
                                </span>
                                <InputBase 
                                class="table-input"
                                placeholder="role" 
                                type="list"
                                :items="roleList"
                                v-model="userRequest.role"/>
                                <span>
                                    <b>💸 balance:</b>
                                </span>
                                <InputBase 
                                class="table-input"
                                placeholder="balance" 
                                v-model="userRequest.balance"/>
                            </div>
                        </td>
                        <td 
                        colspan="1"
                        class="action-cell">
                            <div class="action-wrapper confirm-button">
                                <ButtonBase
                                    @click="handleCreateUser"
                                    :icon="confirmIcon"
                                />
                            </div>
                        </td>
                    </template>
                </tr>
            </table>
        </div>
        <div class="button-content">
            <div class="button-wrapper">
                <ButtonBase 
                @click="addUserRow"
                :icon="addIcon"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import ButtonBase from './ButtonBase.vue';
import { onMounted, reactive, ref } from 'vue';
import api from '../common/api';
import { addIcon, confirmIcon, removeIcon, userIcon } from '../main';
import InputBase from './InputBase.vue';

const roleList = ["user", "admin"]

const users = ref([])
const userRequest = reactive({
    "id": 0,
    "username": "",
    "password": "",
    "balance": null,
    "role": "user"
})

const handleAllUsers = async() => {
    try {
        const response = await api.get("/admin/all-users")
        users.value = response.data
        console.log(response.data)
    } catch(e) {
        console.error(e)
    }
}

const addUserRow = () => {
    const alreadyAdding = users.value.some(u => u.id === 0);
    
    if (!alreadyAdding) {
        users.value.push(userRequest);
    } else {
        alert("You are already adding a user! Fill out the current row first. ✋");
    }
}

const handleCreateUser = async() => {
    try {
        console.log(userRequest)
        const response = await api.post("/admin/create-user", {
            "username": userRequest.username,
            "password": userRequest.password,
            "balance": userRequest.balance,
            "role": userRequest.role,
        })

        if (response.status === 200 || response.status === 201) {
            await handleAllUsers();
            
            Object.assign(userRequest, {
                id: 0,
                username: "",
                password: "",
                balance: null,
                role: "user"
            });
        }
    } catch(e){
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
    gap: 0.5rem;
}

.table-content {
    display: flex;
    justify-content: center;
    width: 100%;
}

.button-content {
    width: 80%;
    display: flex;
    justify-content: flex-start;
}

.confirm-button :deep(button) {
    padding: 16px;
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
    height: 100%;
}

.action-cell {
    text-align: center;
}

.action-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.25rem;
    width: 100%;
    height: 100%;
}

.confirm-button :deep(button){
    padding: 8px;
}

td .input-container span {
    display: flex;
    align-self: flex-start;
    width: 65%;
}

td .input-container{
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.4rem;
    width: 100%;
}

td .table-input {
    width: 100%; 
}

td .table-input :deep(input) {
    padding: 8px;
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

@media (max-width: 768px) {
    table, .button-content {
        width: 100%;
    }

    table {
        min-width: 600px; 
    }

    .table-content {
        width: 100%;
        max-width: 100vw;
        overflow-x: auto; 
        justify-content: flex-start;
        padding-bottom: 0.5rem; 
    }

    tr td {
        padding: 0.4rem 0.5rem;
    }
}
</style>