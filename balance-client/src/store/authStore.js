import { defineStore } from 'pinia' 
import { ref } from 'vue'
 
export const authStore = defineStore('auth_store', () => {
    const isAdmin = ref(bool)
    const isValid = ref(bool)
    const isActive = ref(bool)

    return {
        isAdmin,
        isValid, 
        isActive
    }
})