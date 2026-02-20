import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import { router } from './common/router'

export { default as addIcon } from "./assets/add.svg";
export { default as removeIcon } from "./assets/remove.svg" 
export { default as userIcon } from "./assets/user.svg"

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.mount('#app')
