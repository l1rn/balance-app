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

router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore()
    
    if (to.path === '/login') {
        if(authStore.isAuth){
            return next('/main')
        }
        return next()
    } 
    await authStore.checkAuth()
    
    if(to.meta.requiresAuth && !authStore.isAuth){
        next('/login')
    } else {
        next()
    }

    if(to.path === '/login') {
        return next()
    }
})