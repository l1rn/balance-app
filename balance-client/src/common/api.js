import axios from "axios";
import { useAuthStore } from "../store/authStore";

const api = axios.create({
    baseURL: `${import.meta.env.VITE_API_URL}/api`,
    withCredentials: true
})

api.interceptors.response.use(
    async (response) => {
        if(response.status === 200 && window.location.pathname === "/login"){
            window.location.href = "/main"
        } 
        return response
    },
    (error) => {
        if(error.response && error.response.status === 401){
            const authStore = useAuthStore()
            document.cookie = "token=null; max-age=0; path=/;"
            authStore.setAutorization(false)
        } else if(error.response && error.response.status === 403) {
            alert("Access denied.")
        }
        return Promise.reject(error)
    }
)

export default api 