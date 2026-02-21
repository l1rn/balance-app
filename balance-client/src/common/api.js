import axios from "axios";
import { useAuthStore } from "../store/authStore";
import { router } from "./router";

const api = axios.create({
    baseURL: `${import.meta.env.VITE_API_URL}/api`,
    withCredentials: true
})

api.interceptors.response.use(
    async (response) => {
        if(response.status === 200 && window.location.pathname === "/login"){
            router.push("/main")
        } 
        return response
    },
    (error) => {
        if(
            error.response && 
            error.response.status === 401 &&
            router.currentRoute.value.path !== "login"){
            const authStore = useAuthStore()
            document.cookie = "token=null; max-age=0; path=/;"
            authStore.setAuthorization(false)
            router.replace("/login")
        } else if(error.response && error.response.status === 403) {
            alert("Access denied.")
        }
        return Promise.reject(error)
    }
)

export default api 