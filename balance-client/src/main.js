import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import { router } from './common/router'

export { default as addIcon } from "./assets/add.svg";
export { default as removeIcon } from "./assets/remove.svg" 
export { default as userIcon } from "./assets/user.svg"
export { default as confirmIcon } from "./assets/confirm.svg"
export { default as logoutIcon } from "./assets/logout.svg"

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.mount('#app')
