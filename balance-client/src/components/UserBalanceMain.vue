<template>
    <div class="main-wrapper">
        <div class="main-container">
            <div class="info-container">
                <div class="id">
                    ID: {{ user.id }}
                </div>
                <div class="username">
                    User: {{ user.username }}
                </div>
                <div class="role">
                    Role: {{ user.role }}
                </div>
                <div class="action-container">
                    <div class="button-group">
                        <button 
                        @click="setMode('plus')"
                        :class="{ active: currentMode === 'plus'}">
                            plus
                        </button>
                        <button 
                        @click="setMode('minus')"
                        :class="{ active: currentMode === 'minus'}">
                            minus
                        </button>
                        <button>edit</button>
                    </div>
                    <div v-if="currentMode" class="amount-container">
                        <input type="text" v-model="amount">
                        <button @click="handleConfirm" class="confirm-btn">
                            {{ currentMode === 'plus' ? 'Add Funds' : 'Withdraw' }}
                        </button>
                    </div>
                </div>
            </div>
            <div class="balance-container">
                <div class="text">
                    {{ user.balance }} ₹
                </div>
            </div>
        </div>
        <div v-if="user.transactions" class="table-container">
            <table>
                <tr>
                    <td>Username</td>
                    <td>Amount</td>
                    <td>Type</td>
                    <td>Created At</td>
                </tr>
                <tr 
                v-for="t in user.transactions">
                    <td>{{ t.username }}</td>
                    <td>{{ t.amount }}</td>
                    <td>{{ t.type }}</td>
                    <td class="time-cell">
                        {{ t.created_at.split("T")[1].split(".")[0] }}
                        <div class="date">
                            {{ t.created_at.split("T")[0] }}
                        </div>
                    </td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';
import api from '../common/api';
import { useRoute} from "vue-router"

const currentMode = ref('')
const amount = ref('')

const setMode = (mode) => {
    if(currentMode.value === mode){
        currentMode.value = ''
    } else {
        currentMode.value = mode
    }
}

const route = useRoute() 
const user = ref({
    id: "",
    username: "",
    role: "",
    balance: "",
    transactions: [],
})

const transactionsHandle = async() => {
    try {
        const response = await api.get(`/admin/users/${route.params.id}`)
        user.value = response.data
    } catch(e){
        console.error(e)
    }
}

const topupHandle = async() => {
    try {
        const response = await api.post('/admin/top-up', {
            user_id: user.value.id,
            amount: Number(amount.value),
        })
    } catch(e){
        console.error(e)
    }
}

const withdrawHandle = async() => {
    try {
        const response = await api.post('/admin/withdraw', {
            user_id: user.value.id,
            amount: Number(amount.value),
        })
    } catch(e){
        console.error(e)
    }
}

const handleConfirm = async() => {
    if(amount.value <= 0) return alert("Enter a valid number")
    if(currentMode.value === 'plus') await topupHandle()
    else if(currentMode.value === 'minus') await withdrawHandle()

    amount.value = 0
    currentMode.value = ''
    await transactionsHandle()
}

onMounted(async() => {
    await transactionsHandle()
})

watch(() => route.params.id, () => {
    transactionsHandle()
})
</script>

<style scoped>
.main-wrapper {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    gap: 1.5rem;
}

.main-container {
    width: 70%;
    display: flex;
    justify-content: center;
    gap: 2rem;

}

.info-container {
    display: flex;
    flex-direction: column;
    width: 60%;
    font-size: 32px;
}

.balance-container {
    border-radius: 32px;
    width: 450px;
    height: 230px;
    padding: 2rem;
    background: #202020;
    display: flex;
    justify-content: center;
    align-items: center;
}

.balance-container .text{
    font-size: 96px;
}

.table-container {
    display: flex;
    width: 70%;
}

table {
    width: 100%;
    border-spacing: 0;
    border-collapse: separate;
}

.table-container tr td {
    padding: 0.5rem 1rem;
    border: 0;
    background: #2a2a2a;
    text-align: center;
}

.time-cell {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
}

.time-cell .date {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 13px;
    background: #323131;
    border-radius: 8px;
    padding: 0.1rem 0;
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

@media only screen and (max-width: 600px) {
    .main-container{
        display: flex;
        flex-direction: column;
    }
    .balance-container .text {
        font-size: 48px;
    }
    .balance-container {
        width: 225px;
        height: 115px;
    }
    .info-container {
        font-size: 24px;
    }

    .table-container {
        width: 95%;
        
    }
    .table-container tr td {
        padding: 0.5rem;
    }
}
</style>