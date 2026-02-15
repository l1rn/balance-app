import LoginModal from "../components/LoginModal.vue";
import UserBalanceList from "../components/UserBalanceList.vue";
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from "../store/authStore";

const routes = [
    { 
        path: '/main', 
        component: UserBalanceList,
        meta: { requiresAuth: true },
    },
    { path: '/login', component: LoginModal },
    { path: '/', redirect: '/main' },
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore()
    await authStore.checkAuth()
    if(to.meta.requiresAuth && !authStore.isAuth){
        next('/login')
    } else if (to.path === '/login' && authStore.isAuth) {
        next('/main')
    } else {
        next()
    }
})