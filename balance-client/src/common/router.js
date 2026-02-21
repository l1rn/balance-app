import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from "../store/authStore";

const routes = [
    { 
        path: '/main', 
        component: () => import("../components/UserBalanceList.vue"),
        meta: { requiresAuth: true },
    },
    { path: '/', redirect: '/main' },
    { 
        path: '/users/:id', 
        component: () => import("../components/UserBalanceMain.vue"),
    },
    { 
        path: '/login', 
        component: () => import("../components/LoginModal.vue")
    },
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach(async (to, from) => {
    const authStore = useAuthStore()
    
    console.log("Navigating to:", to.path)
    if(authStore.isAuthenticated === null) {
        await authStore.checkAuth()
    }

    if(to.meta.requiresAuth && !authStore.isAuthenticated){
        return '/login'
    }

    if (to.path === '/login' && authStore.isAuthenticated) {
        return '/main'
    }

    return true
})