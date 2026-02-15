import { defineStore } from 'pinia' 
import { ref} from 'vue'
import api from '../api'
 
export const useAuthStore = defineStore('auth_store', () => {
    const isAdmin = ref(false)
    const isAuth = ref(false)

    const confirmAuth = () => {
        isAuth.value = true;
    }

    const confirmLougout = () => {
        isAuth.value = false;
    }

    const checkAuth = async() => {
        const r = await api.get("/admin/check-balance/1");
        if(r.status === 200){
            confirmAuth()
            return true
        } else {
            confirmLougout()
            return false
        }
    }

    return {
        isAdmin,
        isAuth,
        confirmAuth,
        confirmLougout,
        checkAuth
    }
})