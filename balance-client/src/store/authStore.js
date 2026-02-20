import { defineStore } from 'pinia' 
import { ref} from 'vue'
import api from '../common/api'
 
export const useAuthStore = defineStore('auth_store', () => {
    const isAdmin = ref(false)
    const isAuthenticated = ref(null)

    const setAuthorization = (value) => {
        isAuthenticated.value = value;
    }

    const checkAuth = async() => {
        const r = await api.get("/admin/check-balance/1");
        if(r.status === 200){
            setAuthorization(true)
        } else {
            setAuthorization(false)
        }
    }

    return {
        isAdmin,
        isAuthenticated,
        setAuthorization,
        checkAuth
    }
})